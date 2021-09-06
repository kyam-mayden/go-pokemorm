package models

import (
	"gorm.io/gorm"
	"time"
)

func (Trainer) TableName() string {
	return "trainer"
}

type Trainer struct {
	gorm.Model
	ID               uint `gorm:"primaryKey"`
	FirstName        string `validate:"required,min=1,max=256"`
	SecondName       string `validate:"required,min=1,max=256"`
	HomeTown         int64 `validate:"required"`
	FavouriteType    int64 `validate:"required"`
	FavouritePokemon int64 `validate:"required"`
	Evil             bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt

	// Relationships
	HomeTownLocation Location `gorm:"foreignKey:HomeTown"`
}

func init() {
	Connect()
	db = GetDB()
	db.AutoMigrate(&Trainer{})
}

func (b *Trainer) CreateTrainer() (*Trainer , *gorm.DB) {
	db := db.Create(&b)
	return b, db
}

func GetAllTrainers() []Trainer {
	var Trainers []Trainer
	db.Preload("HomeTownLocation").Find(&Trainers)
	return Trainers
}

func GetTrainerById(Id int64) (*Trainer , *gorm.DB) {
	var getTrainer Trainer
	db := db.Where("id = ?", Id).First(&getTrainer)
	return &getTrainer, db
}

func DeleteTrainer(Id int64) Trainer {
	var book Trainer
	db.Where("id = ?", Id).Delete(book)
	return book
}
