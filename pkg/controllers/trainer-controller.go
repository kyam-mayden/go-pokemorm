package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/pkg/models"
	"net/http"
	"strconv"
)

func CreateTrainer(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
    newTrainer := &models.Trainer{}

    dataError := json.NewDecoder(request.Body).Decode(&newTrainer)
    if dataError != nil {
        http.Error(responseWriter, dataError.Error(), http.StatusBadRequest)
        fmt.Println("Error when creating new trainer: ", dataError.Error())

        return
    }

    if !ValidateTrainerCreateRequest(request, newTrainer) {
        responseWriter.WriteHeader(http.StatusBadRequest)

        return
    }

	b, db:= newTrainer.CreateTrainer()

	if db.Error != nil {
		fmt.Println("Error when creating new trainer: ", db.Error)
		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	data,_ := json.Marshal(b)
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(data)
}

func ValidateTrainerCreateRequest(request *http.Request, newTrainer *models.Trainer) bool {
    validator := validator.New()

    err := validator.Struct(newTrainer)

    if err != nil {
        fmt.Println("Validation error when creating new trainer: ", err.Error())

        return false
    }

    return true
}

func GetTrainer(w http.ResponseWriter, r *http.Request) {
	newTrainers:= models.GetAllTrainers()
	data, _ := json.Marshal(newTrainers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetTrainerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	trainerId := vars["trainerId"]
	ID, err:= strconv.ParseInt(trainerId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	trainerDetails, db:= models.GetTrainerById(ID)
	if db.Error != nil {
		fmt.Println("Record not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(trainerDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
//
//func UpdateTrainer(w http.ResponseWriter, r *http.Request) {
//	var updateTrainer = &models.Trainer{}
//	utils.ParseBody(r, updateTrainer)
//	vars := mux.Vars(r)
//	trainerId := vars["trainerId"]
//	ID, err:= strconv.ParseInt(trainerId, 0, 0)
//	if err != nil {
//		fmt.Println("Error while parsing")
//	}
//	trainerDetails, db:= models.GetTrainerById(ID)
//	if updateTrainer.Name != "" {
//		trainerDetails.Name = updateTrainer.Name
//	}
//	if updateTrainer.Author != "" {
//		trainerDetails.Author = updateTrainer.Author
//	}
//	if updateTrainer.Publication != "" {
//		trainerDetails.Publication = updateTrainer.Publication
//	}
//	db.Save(&trainerDetails)
//	res, _ := json.Marshal(trainerDetails)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}
//
//func DeleteTrainer(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	trainerId := vars["trainerId"]
//	ID, err:= strconv.ParseInt(trainerId, 0, 0)
//	if err != nil {
//		fmt.Println("Error while parsing")
//	}
//	trainer:= models.DeleteTrainer(ID)
//	res, _ := json.Marshal(trainer)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}
