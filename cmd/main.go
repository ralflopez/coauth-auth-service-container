package main

import (
	db "coauth/pkg/config/db"
	di "coauth/pkg/config/di"
	server "coauth/pkg/config/server"
	"coauth/pkg/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Database
	db, queries, err := db.StartDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Server
	r := mux.NewRouter()
	l := log.New(os.Stdout, "[User API] ", log.LstdFlags)
	s := server.NewServer(r, queries, l)
	di := di.NewDIContainer(s)

	// Routes
	routes.RegisterRoutes(s, di)

	// Server: Start
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), s.Router)
	if err == nil {
		return err
	}

	return nil
}