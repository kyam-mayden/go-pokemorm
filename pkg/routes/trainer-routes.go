package routes

import (
	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/pkg/controllers"
	"github.com/marty-crane/go-pokemorm/pkg/middleware"
)

var RegisterTrainerStoreRoutes = func(router *mux.Router) {
	router.Handle("/trainer", middleware.OAuthMiddleware(controllers.GetTrainer)).Methods("GET")
	router.Handle("/trainer", middleware.BasicAuthMiddleware(controllers.CreateTrainer)).Methods("POST")
	router.Handle("/trainer/{trainerId}", middleware.BasicAuthMiddleware(controllers.GetTrainerById)).Methods("GET")

	//router.HandleFunc("/trainer/{trainerId}", controllers.UpdateTrainer).Methods("PUT")
	//router.HandleFunc("/trainer/{trainerId}", controllers.DeleteTrainer).Methods("DELETE")
}
