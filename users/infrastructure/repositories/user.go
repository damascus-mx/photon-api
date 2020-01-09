package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/damascus-mx/photon-api/users/entity"
	_ "github.com/lib/pq"
)

// UserRepository Handles all persistence user operations
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository Returns user repository instance
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// ---- USER OPERATIONS ----

// Save Inserts user into persistence layer
func (u *UserRepository) Save(user *entity.UserModel) (int, error) {
	// Store new object into DB
	statement := `INSERT INTO users (name, surname, birth, username, password) VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	id := 0
	err := u.DB.QueryRow(statement, user.Name, user.Surname, user.Birth, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	} else if id == 0 {
		return 0, errors.New("Cannot save user")
	}

	return id, nil
}

// FetchByID Get user by ID
func (u *UserRepository) FetchByID(id int64) (*entity.UserModel, error) {
	statement := `SELECT * FROM users WHERE id = $1`
	user := new(entity.UserModel)
	err := u.DB.QueryRow(statement, id).Scan(&user.ID, &user.Name, &user.Surname, &user.Birth, &user.Username, &user.Password, &user.Image,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

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

	return nil
}

// FetchByUsername Retrieves a user by username
func (u *UserRepository) FetchByUsername(username string) (*entity.UserModel, error) {
	statement := `SELECT * FROM users WHERE username = $1`
	user := new(entity.UserModel)
	err := u.DB.QueryRow(statement, username).Scan(&user.ID, &user.Name, &user.Surname, &user.Birth, &user.Username, &user.Password, &user.Image,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
