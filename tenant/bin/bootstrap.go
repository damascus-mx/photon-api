package bin

import (
	"log"
	"os"

	"github.com/damascus-mx/photon-api/tenant/core/config"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

// InitHTTPServer Initialize a new HTTP Server
func InitHTTPServer() *chi.Mux {
	router := chi.NewRouter()
	initEnvironment()
	return router
}

func initEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	config.DBConnection = os.Getenv("DB_CONNECTION")
}
