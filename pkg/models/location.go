package models

import (
	"github.com/marty-crane/go-pokemorm/pkg/config"
	"gorm.io/gorm"
)

func (Location) TableName() string {
	return "location"
}

type Location struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	Name   string
	Region string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Location{})
}
