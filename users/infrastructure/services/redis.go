package infrastructure

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

// InitRedis Initialize Redis client
func InitRedis(address string, password string) (c *redis.Client) {
	c = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	pong, err := c.Ping().Result()
	fmt.Println(pong, err)
	return c
}
