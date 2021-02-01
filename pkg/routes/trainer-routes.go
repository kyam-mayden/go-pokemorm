package routes

import (
	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/pkg/controllers"
)

var RegisterTrainerStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/trainer", controllers.CreateTrainer).Methods("POST")
	router.HandleFunc("/trainer", controllers.GetTrainer).Methods("GET")
	router.HandleFunc("/trainer/{trainerId}", controllers.GetTrainerById).Methods("GET")
	//router.HandleFunc("/trainer/{trainerId}", controllers.UpdateTrainer).Methods("PUT")
	//router.HandleFunc("/trainer/{trainerId}", controllers.DeleteTrainer).Methods("DELETE")
}
