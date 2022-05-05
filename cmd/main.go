package main

import (
	"coauth/pkg/config"
	"coauth/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Server
	r := mux.NewRouter()

	// Routes
	routes.RegisterUserRoutes(r)
	
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), r)
	if err == nil {
		log.Fatal(err.Error())
	}
}