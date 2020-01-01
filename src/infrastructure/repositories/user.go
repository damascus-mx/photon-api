package infrastructure

import (
	"database/sql"
	"fmt"

	entity "github.com/damascus-mx/photon-api/src/entity"
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

// Save Inserts user into persistence layer
func (u *UserRepository) Save(user *entity.UserModel) error {
	// Store new object into DB
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return err
	}
	fmt.Println(rows)

	defer rows.Close()

	return nil
}

// FetchByID Get user by ID
func (u *UserRepository) FetchByID(id int) *entity.UserModel {
	/*row := u.DB.QueryRow("SELECT * FROM users WHERE id = %i", id)
	if row == nil {
		return nil
	}

	row.*/
	return nil
}

// FetchAll Get all users
func (u *UserRepository) FetchAll() ([]*entity.UserModel, error) {
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*entity.UserModel, 0)
	for rows.Next() {
		user := new(entity.UserModel)
		fmt.Println(user)
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
