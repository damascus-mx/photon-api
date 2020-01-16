package service

import (
	"fmt"

	env "github.com/damascus-mx/photon-api/authentication/common/config"
	"github.com/go-redis/redis/v7"
)

// InitRedis Get a new redis client pool
func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_CONNECTION,
		Password: env.REDIS_PASSWORD,
		DB:       0,
	})

	conn := client.Conn()
	defer conn.Close()

	pong, err := conn.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\nRedis ping - %s\n", pong)
	return client
}
