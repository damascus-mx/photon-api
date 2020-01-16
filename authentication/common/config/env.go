package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_CONNECTION    string
	MQ_CONNECTION    string
	REDIS_CONNECTION string
	REDIS_PASSWORD   string
	JWT_SECRET       string
	PHOTON_SECRET    string
)

// StartEnvironment Intilialize Environment variables
func StartEnvironment() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	DB_CONNECTION = os.Getenv("DB_CONNECTION")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	MQ_CONNECTION = os.Getenv("MQ_CONNECTION")
	PHOTON_SECRET = os.Getenv("PHOTON_SECRET")
	REDIS_CONNECTION = os.Getenv("REDIS_CONNECTION")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
}
