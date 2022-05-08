package config

import (
	"coauth/pkg/db"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, *db.Queries, error) {
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")
	database, err := sql.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", user, password, name))
	if err != nil {
		return nil, nil, err
	}

	queries := db.New(database)
	return database, queries, nil
}