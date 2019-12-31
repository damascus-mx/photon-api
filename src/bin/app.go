package app

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	common "github.com/damascus-mx/photon-api/src/core/common"
)

// InitApplication Creates a new router instance
func InitApplication() *chi.Mux {
	// Create new Router
	fmt.Print("Running Photon REST Microservice\n")

	// Load virtual environment
	initEnvironment()

	// Start router
	var router common.IRouter = &common.Router{}
	mux := router.InitializeRouter(chi.NewRouter())

	// Set routes
	router.SetRoutes(mux)

	// Load Redis client
	redis := common.InitRedis(os.Getenv("REDIS_CONNECTION"), os.Getenv("REDIS_PASSWORD"))
	fmt.Print(redis.ClientID().String())

	// Connect to DB
	db := common.ConnectDB(os.Getenv("DB_CONNECTION"))
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return mux
}

func initEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3BucketKey := os.Getenv("AWS_S3_KEY")

	fmt.Printf("AWS Key %s", s3BucketKey)
}

/*
// setRoutes Attach routes/handlers
func setRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, err := utils.BuildMessage([]byte("Welcome from JSON Method"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something happened"))
		}
		w.Write(res)
	})
}
*/
