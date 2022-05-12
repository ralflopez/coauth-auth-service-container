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
	"github.com/gorilla/sessions"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Env
	// err := godotenv.Load()
	// if err != nil {
	//   log.Fatal("Error loading .env file")
	// }

	// Database
	db, queries, err := db.StartDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Session Store
	var sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	// sessionStorePort := ":" + os.Getenv("REDIS_PORT")
	// sessionStorePassword := os.Getenv("REDIS_PASSWORD")
	// sessionStore, err := redistore.NewRediStore(10, "tcp", sessionStorePort, sessionStorePassword, []byte("secret-key"))
	// if err != nil {
	// 	panic(err)
	// }
	// defer sessionStore.Close()

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
	fmt.Printf("Server listening in port: %v\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), s.Router)
	if err == nil {
		return err
	}

	return nil
}
