package usecase

import (
	"time"

	config "github.com/damascus-mx/photon-api/src/core/config"
	helper "github.com/damascus-mx/photon-api/src/core/helper"
	entity "github.com/damascus-mx/photon-api/src/entity"
)

// UserRepository Persistence user layer
type UserRepository interface {
	Save(user *entity.UserModel) (int, error)
	FetchByID(id int64) (*entity.UserModel, error)
	FetchAll() ([]*entity.UserModel, error)
}

// UserUsecase User usecase
type UserUsecase struct {
	userRepository UserRepository
}

// NewUserUsecase Exports User usecase instance
func NewUserUsecase(userRepository UserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

// ---- USER OPERATIONS ----

// CreateUser Save a new user
func (u *UserUsecase) CreateUser(name, surname, username, password, birth string) (int, error) {
	// Convert birth field to Time
	birthFormatted, err := time.Parse(config.RFC3339Tiny, birth)
	if err != nil {
		return 0, err
	}

	// Map request form to entity
	user := entity.NewUser()
	user.Name = name
	user.Surname = surname
	user.Birth = birthFormatted
	user.Username = username
	user.Password = password

	// Verify user state
	err = user.Validate()
	if err != nil {
		return 0, err
	}

	// Encrypt password and store the retrieved hash
	hash64, err := helper.HashString(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hash64

	// Send sanitized user to repository
	return u.userRepository.Save(user)
}

// GetAllUsers Retrieves all users
func (u *UserUsecase) GetAllUsers() ([]*entity.UserModel, error) {
	return u.userRepository.FetchAll()
}

// GetUserByID Retrieves user by ID or username
func (u *UserUsecase) GetUserByID(id int64) (*entity.UserModel, error) {
	return u.userRepository.FetchByID(id)
}
