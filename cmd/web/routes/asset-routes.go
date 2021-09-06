package routes

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/marty-crane/go-pokemorm/pkg/config"
)

var RegisterAssetRoutes = func(router *mux.Router) {
    assetsDirectory := config.Get("STATIC_ASSETS", ".ui/static")
    fileServer := http.FileServer(http.Dir(assetsDirectory))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))
}
