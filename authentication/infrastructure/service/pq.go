package service

import (
	"database/sql"
	env "github.com/damascus-mx/photon-api/authentication/common/config"
	_ "github.com/lib/pq"
)

// InitPostgres Get a postgres connection pool
func InitPostgres() *sql.DB {
	db, err := sql.Open("postgres", env.DB_CONNECTION)
	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}
