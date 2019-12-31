package core

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB Get a connection to the DB
func ConnectDB(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
