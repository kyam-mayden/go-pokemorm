package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

var RegisterAssetRoutes = func(router *mux.Router) {
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))
}
