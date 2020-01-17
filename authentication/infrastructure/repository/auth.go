package repository

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

// AuthRepository Authentication repository
type AuthRepository struct {
	Cache *redis.Client
}

// NewAuthRepository Get new authentication repository
func NewAuthRepository(cacheClient *redis.Client) *AuthRepository {
	return &AuthRepository{Cache: cacheClient}
}

// GetUserToken Get user's token
func (a *AuthRepository) GetUserToken(username string) (string, error) {
	conn := a.Cache.Conn()
	defer conn.Close()

	token, err := conn.Get(fmt.Sprintf("auth:token:%s", username)).Result()
	if err != nil {
		return "", nil
	}

	return token, nil
}

// SetUserToken Set user's token
func (a *AuthRepository) SetUserToken(username, token string) error {
	conn := a.Cache.Conn()
	defer conn.Close()

	err := conn.Set(fmt.Sprintf("auth:token:%s", username), token, (time.Duration(336) * time.Hour)).Err()
	if err != nil {
		return err
	}

	return nil
}
