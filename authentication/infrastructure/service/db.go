package service

import (
	"database/sql"
	"log"

	env "github.com/damascus-mx/photon-api/authentication/common/config"

	// PQ Driver
	_ "github.com/lib/pq"
)

// InitDatabase Get a postgres connection pool
func InitDatabase() *sql.DB {
	db, err := sql.Open("postgres", env.DB_CONNECTION)
	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	log.Printf(env.ServiceConnected, "Database")
	return db
}
