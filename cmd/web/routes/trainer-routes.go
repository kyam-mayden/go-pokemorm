package routes

import (
	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/cmd/web/handlers"
	"github.com/marty-crane/go-pokemorm/cmd/web/middleware"
)

var RegisterTrainerRoutes = func(router *mux.Router) {
	router.HandleFunc("/", handlers.RenderTrainers).Methods("GET")
	router.Handle("/trainer", middleware.OAuthMiddleware(handlers.GetTrainer)).Methods("GET")
	router.Handle("/trainer", middleware.OAuthMiddleware(handlers.CreateTrainer)).Methods("POST")
	router.Handle("/trainer/{trainerId}", middleware.OAuthMiddleware(handlers.GetTrainerById)).Methods("GET")

	//router.HandleFunc("/trainer/{trainerId}", controllers.UpdateTrainer).Methods("PUT")
	//router.HandleFunc("/trainer/{trainerId}", controllers.DeleteTrainer).Methods("DELETE")
}
