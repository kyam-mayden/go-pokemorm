package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/marty-crane/go-pokemorm/cmd/web/routes"
	"github.com/marty-crane/go-pokemorm/pkg/config"
	"github.com/marty-crane/go-pokemorm/pkg/logger"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTrainerRoutes(r)
	routes.RegisterAssetRoutes(r)
	http.Handle("/", r)

    port := config.Get("WEB_PORT", ":8080")

	logger.Info(fmt.Sprintf("Serving on port %s", port))

	fmt.Println()
	err := http.ListenAndServe(fmt.Sprintf("localhost%s", port), r);

	log.Fatal(err)
	logger.Fatal(fmt.Sprint(err))
}
