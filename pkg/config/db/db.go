package config

import (
	"coauth/pkg/db"
	"database/sql"

	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, *db.Queries, error) {
	database, err := sql.Open("postgres", "user=user password=password dbname=coauth sslmode=disable")
	if err != nil {
		return nil, nil, err
	}

	queries := db.New(database)
	return database, queries, nil
}