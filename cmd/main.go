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
	redisstore "gopkg.in/boj/redistore.v1"
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

	// Session Store
	// var sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	sessionStore, err := redisstore.NewRediStore(10, "tcp", ":63791", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer sessionStore.Close()

	// Server
	r := mux.NewRouter()
	l := log.New(os.Stdout, "[User API] ", log.LstdFlags)
	s := server.NewServer(r, queries, l, sessionStore)
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