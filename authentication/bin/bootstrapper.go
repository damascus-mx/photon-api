package bin

import (
	"database/sql"
	"log"

	env "github.com/damascus-mx/photon-api/authentication/common/config"
	delivery "github.com/damascus-mx/photon-api/authentication/infrastructure/delivery/http"
	"github.com/damascus-mx/photon-api/authentication/infrastructure/service"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v7"
	"github.com/streadway/amqp"
)

type Bootstrapper struct {
	redis *redis.Client
	db    *sql.DB
	mq    *amqp.Connection
}

// StartServices Initialize required services
func (b *Bootstrapper) StartServices() {
	env.StartEnvironment()

	// Init redis service
	b.redis = service.InitRedis()

	// Init Database
	b.db = service.InitPostgres()
}

// StartHTTP Initialize HTTP Service bootstrapping
func (b *Bootstrapper) StartHTTP() *chi.Mux {
	log.Println("Starting Photon's Authentication service HTTP Server")
	router := new(delivery.HTTPRouter)
	mux := router.NewRouter(b.db, b.redis)

	return mux
}
