package models

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"github.com/marty-crane/go-pokemorm/pkg/config"
)

var (
	db * gorm.DB
)

func Connect() {
    dataSource := getDSN()

	d, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dataSource, // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil{
		panic(err)
	}
	db = d
}

func getDSN() string {
    databaseUser := config.Get("DATABASE_USER", "sail")
    databasePassword := config.Get("DATABASE_PASSWORD", "password")
    databaseAddress := config.Get("DATABASE_ADDRESS", "127.0.0.1")
    databasePort := config.Get("DATABASE_PORT", "3306")
    databaseName := config.Get("DATABASE_NAME", "pokemorm")

    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        databaseUser,
        databasePassword,
        databaseAddress,
        databasePort,
        databaseName,
    )
}

func GetDB() *gorm.DB {
	return db
}
