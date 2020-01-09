package app

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	delivery "github.com/damascus-mx/photon-api/users/infrastructure/delivery/http"
	services "github.com/damascus-mx/photon-api/users/infrastructure/services"
)

// InitApplication Creates a new router instance
func InitApplication() *chi.Mux {
	// Create new Router
	fmt.Print("Running Photon REST USER Microservice\n")

	// Load virtual environment
	initEnvironment()

	// ---> LOAD REQUIRED SERVICES <---
	// Load Redis client
	redis := services.InitRedis(os.Getenv("REDIS_CONNECTION"), os.Getenv("REDIS_PASSWORD"))
	fmt.Print(redis.ClientID().String())

	// Load main postgres DB pool
	db, err := services.ConnectPool(os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatal(err)
	}

	// Start router
	var router delivery.IRouter = delivery.NewHTTPRouter(db, redis)
	mux := router.InitializeRouter(chi.NewRouter())

	// Set routes
	router.SetRoutes(mux)

	return mux
}

func initEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
