package models

import (
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
	Connect()
	db = GetDB()
	db.AutoMigrate(&Location{})
}
