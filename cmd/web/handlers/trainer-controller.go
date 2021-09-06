package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/pkg/models"
	"github.com/marty-crane/go-pokemorm/pkg/services"
	"net/http"
	"strconv"
)

func CreateTrainer(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	trainerService := services.TrainerService{}

	model, error := trainerService.CreateTrainer(request)

	if error != nil {
	    responseWriter.WriteHeader(http.StatusBadRequest)

	    return
	}

	data,_ := json.Marshal(model)
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(data)
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
