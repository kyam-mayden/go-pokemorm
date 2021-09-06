package routes

import (
	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/cmd/web/handlers"
	"github.com/marty-crane/go-pokemorm/cmd/web/middleware"
	"net/http"
)

var RegisterTrainerStoreRoutes = func(router *mux.Router) {
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	router.HandleFunc("/", handlers.RenderTrainers).Methods("GET")
	router.Handle("/trainer", middleware.OAuthMiddleware(handlers.GetTrainer)).Methods("GET")
	router.Handle("/trainer", middleware.OAuthMiddleware(handlers.CreateTrainer)).Methods("POST")
	router.Handle("/trainer/{trainerId}", middleware.OAuthMiddleware(handlers.GetTrainerById)).Methods("GET")

	//router.HandleFunc("/trainer/{trainerId}", controllers.UpdateTrainer).Methods("PUT")
	//router.HandleFunc("/trainer/{trainerId}", controllers.DeleteTrainer).Methods("DELETE")
}
