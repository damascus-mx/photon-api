package service

import (
	"log"

	env "github.com/damascus-mx/photon-api/authentication/common/config"
	"github.com/go-redis/redis/v7"
)

// InitCache Get a new redis client pool
func InitCache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_CONNECTION,
		Password: env.REDIS_PASSWORD,
		DB:       0,
	})

	conn := client.Conn()
	defer conn.Close()

	_, err := conn.Ping().Result()
	if err != nil {
		panic(err.Error())
	}

	log.Printf(env.ServiceConnected, "Cache")
	return client
}
