package services

import (
	"encoding/json"
	"fmt"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/marty-crane/go-pokemorm/pkg/models"
	"net/http"
)

type TrainerService struct {}

func (b *TrainerService) CreateTrainer(request *http.Request) (*models.Trainer, error) {
    newTrainer := &models.Trainer{}

    dataError := json.NewDecoder(request.Body).Decode(&newTrainer)
    if dataError != nil {
        fmt.Println("Error when creating new trainer: ", dataError.Error())

        return newTrainer, errors.New("Error when creating new trainer: " + dataError.Error())
    }

    if !validateTrainerCreateRequest(request, newTrainer) {

        return newTrainer, errors.New("Validation error when creating new trainer")
    }

	model, db:= newTrainer.CreateTrainer()

	if db.Error != nil {
		fmt.Println("Error when creating new trainer: ", db.Error)

        return newTrainer, errors.New("Error when creating new trainer")
	}

	return model, nil
}

func validateTrainerCreateRequest(request *http.Request, newTrainer *models.Trainer) bool {
    validator := validator.New()

    err := validator.Struct(newTrainer)

    if err != nil {
        fmt.Println("Validation error when creating new trainer: ", err.Error())

        return false
    }

    return true
}
