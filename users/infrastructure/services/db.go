package infrastructure

import (
	"database/sql"
	"log"

	// Postgres Go driver
	_ "github.com/lib/pq"
)

// ConnectPool Get a connection to the Pool
func ConnectPool(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
