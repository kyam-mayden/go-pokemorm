package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/pkg/models"
	"net/http"
	"strconv"
)

//var NewTrainer models.Trainer

//func CreateTrainer(w http.ResponseWriter, r *http.Request) {
//	CreateTrainer := &models.Trainer{}
//	utils.ParseBody(r, CreateTrainer)
//	b:= CreateTrainer.CreateTrainer()
//	res,_ := json.Marshal(b)
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}

func GetTrainer(w http.ResponseWriter, r *http.Request) {
	newTrainers:= models.GetAllTrainers()
	data, _ := json.Marshal(newTrainers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetTrainerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trainerId := vars["trainerId"]
	ID, err:= strconv.ParseInt(trainerId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	trainerDetails, _:= models.GetTrainerById(ID)
	res, _ := json.Marshal(trainerDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
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
