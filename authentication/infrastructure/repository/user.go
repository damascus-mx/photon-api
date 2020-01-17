package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/damascus-mx/photon-api/authentication/entity"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

// UserRepository User repository
type UserRepository struct {
	DB    *sql.DB
	Cache *redis.Client
}

// NewUserRepository Get user repository
func NewUserRepository(database *sql.DB, cacheClient *redis.Client) *UserRepository {
	return &UserRepository{DB: database, Cache: cacheClient}
}

// FetchByUsername Get user by username
func (u *UserRepository) FetchByUsername(username string) (*entity.UserModel, error) {
	// Check cache layer
	conn := u.Cache.Conn()
	defer conn.Close()

	user := new(entity.UserModel)

	usrJSON, err := u.Cache.Get(fmt.Sprintf("auth:user:%s", username)).Result()
	if err != nil {
		return nil, err
	} else if usrJSON != "" {
		json.Unmarshal([]byte(usrJSON), &user)
		err = user.Validate()
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	// Search in DB
	statement := `SELECT * FROM users WHERE username = $1 LIMIT 1`
	err = u.DB.QueryRow(statement, username).Scan(&user.ID, &user.Name, &user.Surname, &user.Birth, &user.Username, &user.Password, &user.Image,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	// Send user to cache
	go func() {
		jsonUsr, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = conn.Set(fmt.Sprintf("auth:user:%s", username), jsonUsr, (time.Duration(336) * time.Hour)).Err()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	return user, nil
}
