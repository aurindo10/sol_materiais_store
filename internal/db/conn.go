package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	connStr := "postgresql://docker:sol123@localhost:5432/sol_store?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
