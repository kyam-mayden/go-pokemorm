package models

import (
	"github.com/marty-crane/go-pokemorm/pkg/config"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func (Trainer) TableName() string {
	return "trainer"
}

type Trainer struct {
	gorm.Model
	id               uint `gorm:"primaryKey"`
	FirstName         string
	SecondName       string
	FavouriteType    int64
	FavouritePokemon int64
	Evil             bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Trainer{})
}

func (b *Trainer) CreateTrainer() *Trainer {
	//db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTrainers() []Trainer {
	var Trainers []Trainer
	db.Find(&Trainers)
	return Trainers
}

func GetTrainerById(Id int64) (*Trainer , *gorm.DB){
	var getTrainer Trainer
	db := db.Where("id = ?", Id).First(&getTrainer)
	return &getTrainer, db
}

func DeleteTrainer(Id int64) Trainer {
	var book Trainer
	db.Where("id = ?", Id).Delete(book)
	return book
}
