package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marty-crane/go-pokemorm/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTrainerStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
