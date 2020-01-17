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

// Bootstrapper Start all required services/delivery
type Bootstrapper struct {
	cache *redis.Client
	db    *sql.DB
	mq    *amqp.Connection
}

// StartServices Initialize required services
func (b *Bootstrapper) StartServices() {

	// Init Environment
	env.StartEnvironment()

	// Init Cache service
	b.cache = service.InitCache()

	// Init Database
	b.db = service.InitDatabase()

	// Init MQ Broker
	b.mq = service.InitMQBroker()

	log.Println("SERVICE: All services started")
}

// StartHTTP Initialize HTTP Service bootstrapping
func (b *Bootstrapper) StartHTTP() *chi.Mux {
	log.Println("Starting Photon's Authentication service HTTP Server")
	router := new(delivery.HTTPRouter)
	mux := router.NewRouter(b.db, b.cache)

	return mux
}
