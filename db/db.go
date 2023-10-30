package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	dbURL := "user=mekan password=123 dbname=bank sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil
	}
	return db
}
