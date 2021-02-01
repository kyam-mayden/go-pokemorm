package routes

import (
	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/pkg/controllers"
)

var RegisterTrainerStoreRoutes = func(router *mux.Router) {
	//router.HandleFunc("/book/", controllers.CreateTrainer).Methods("POST")
	router.HandleFunc("/trainer", controllers.GetTrainer).Methods("GET")
	router.HandleFunc("/trainer/{trainerId}", controllers.GetTrainerById).Methods("GET")
	//router.HandleFunc("/book/{bookId}", controllers.UpdateTrainer).Methods("PUT")
	//router.HandleFunc("/book/{bookId}", controllers.DeleteTrainer).Methods("DELETE")
}
