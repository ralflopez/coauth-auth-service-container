package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Config
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Server
	router := mux.NewRouter()
	
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err == nil {
		log.Fatal(err.Error())
	}
}