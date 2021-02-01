package config

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "homestead:secret@tcp(192.168.10.10:3306)/pokemorm?charset=utf8&parseTime=True&loc=Local", // data source name
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

func GetDB() *gorm.DB {
	return db
}
