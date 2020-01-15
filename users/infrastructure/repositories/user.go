package infrastructure

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/damascus-mx/photon-api/users/entity"
	"github.com/go-redis/redis/v7"

	// PSQL Driver
	_ "github.com/lib/pq"
)

// UserRepository Handles all persistence user operations
type UserRepository struct {
	DB    *sql.DB
	Redis *redis.Client
}

// NewUserRepository Returns user repository instance
func NewUserRepository(db *sql.DB, redis *redis.Client) *UserRepository {
	return &UserRepository{db, redis}
}

// ---- USER OPERATIONS ----

// Save Inserts user into persistence layer
func (u *UserRepository) Save(user *entity.UserModel) (int, error) {
	// Store new object into DB
	statement := `INSERT INTO users (name, surname, birth, username, password, role) VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	id := 0
	err := u.DB.QueryRow(statement, user.Name, user.Surname, user.Birth, user.Username, user.Password, user.Role).Scan(&id)
	if err != nil {
		return 0, err
	} else if id == 0 {
		return 0, errors.New("Cannot save user")
	}

	return id, nil
}

// FetchByID Get user by ID
func (u *UserRepository) FetchByID(id int64) (*entity.UserModel, error) {
	// Check cache first, Cache Aside -lazy- pattern
	conn := u.Redis.Conn()
	defer conn.Close()

	user := new(entity.UserModel)
	usrJSON, err := conn.Get(fmt.Sprintf("user:%d", id)).Result()

	if err != nil || usrJSON == "" {
		// Read DB
		statement := `SELECT * FROM users WHERE id = $1`

		err = u.DB.QueryRow(statement, id).Scan(&user.ID, &user.Name, &user.Surname, &user.Birth, &user.Username, &user.Password, &user.Image,
			&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Store data in-memory cache, expires in 2 weeks
		json, err := json.Marshal(user)
		if err != nil {
			return nil, err
		}

		err = conn.Set(fmt.Sprintf("user:%d", id), json, (time.Hour * time.Duration(336))).Err()
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	err = json.Unmarshal([]byte(usrJSON), &user)
	if err != nil {
		return nil, errors.New("Corrupted cache")
	}
	fmt.Printf("\nUser in redis: %v\n", user)

	return user, nil
}

// FetchAll Get all users
func (u *UserRepository) FetchAll(limit, index int64) ([]*entity.UserModel, error) {
	rows, err := u.DB.Query(fmt.Sprintf(`SELECT * FROM users WHERE id > %d ORDER BY id ASC FETCH FIRST %d ROWS ONLY`, index, limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*entity.UserModel, 0)
	for rows.Next() {
		user := new(entity.UserModel)
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Birth, &user.Username, &user.Password, &user.Image,
			&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Delete Removes given user
func (u *UserRepository) Delete(id int64) error {
	_, err := u.DB.Exec(fmt.Sprintf(`DELETE FROM users WHERE id = %d`, id))
	if err != nil {
		return err
	}
	// If deleted, invalidate key in cache layer
	conn := u.Redis.Conn()
	defer conn.Close()
	err = conn.Del(fmt.Sprintf("user:%d", id)).Err()
	if err != nil {
		log.Println("REDIS: Key not found")
	}

	return nil
}

// Update Updates given user
func (u *UserRepository) Update(user *entity.UserModel) error {
	statement := `UPDATE users SET name = $1, surname = $2, birth = $3, username = $4, password = $5, 
	image = $6, role = $7, active = $8, updated_at = $9 WHERE id = $10`
	_, err := u.DB.Exec(statement, user.Name, user.Surname, user.Birth, user.Username, user.Password, user.Image,
		user.Role, user.Active, user.UpdatedAt, user.ID)

	if err != nil {
		return err
	}
	// If updated, apply write-through cache pattern
	conn := u.Redis.Conn()
	defer conn.Close()
	json, err := json.Marshal(user)
	err = conn.Set(fmt.Sprintf("user:%d", user.ID), json, (time.Hour * time.Duration(336))).Err()
	if err != nil {
		log.Println("REDIS: Key not found")
	}

	return nil
}

// FetchByUsername Retrieves a user by username
func (u *UserRepository) FetchByUsername(username string) (*entity.UserModel, error) {
	statement := `SELECT * FROM users WHERE username = $1 LIMIT 1`
	user := new(entity.UserModel)
	err := u.DB.QueryRow(statement, username).Scan(&user.ID, &user.Name, &user.Surname, &user.Birth, &user.Username, &user.Password, &user.Image,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetPasswordHash Get a password hashed by username
func (u *UserRepository) GetPasswordHash(id int64) (string, error) {
	statement := `SELECT password FROM users WHERE id = $1`
	hash := ""
	err := u.DB.QueryRow(statement, id).Scan(&hash)
	if err != nil {
		return "", err
	}

	return hash, nil
}
