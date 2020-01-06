package usecase

import (
	"errors"
	"net/url"
	"time"

	config "github.com/damascus-mx/photon-api/src/core/config"
	helper "github.com/damascus-mx/photon-api/src/core/helper"
	entity "github.com/damascus-mx/photon-api/src/entity"
)

// UserRepository Persistence user layer
type UserRepository interface {
	Save(user *entity.UserModel) (int, error)
	FetchByID(id int64) (*entity.UserModel, error)
	FetchAll(limit, index int64) ([]*entity.UserModel, error)
	Delete(id int64) error
	Update(user *entity.UserModel) error
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
	birthFormatted, err := time.Parse(config.MonthDayYear, birth)
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
func (u *UserUsecase) GetAllUsers(limit, index int64) ([]*entity.UserModel, error) {
	return u.userRepository.FetchAll(limit, index)
}

// GetUserByID Retrieves user by ID or username
func (u *UserUsecase) GetUserByID(id int64) (*entity.UserModel, error) {
	return u.userRepository.FetchByID(id)
}

// DeleteUser Deletes given user row
func (u *UserUsecase) DeleteUser(id int64) error {
	return u.userRepository.Delete(id)
}

// UpdateUser Update given user row
func (u *UserUsecase) UpdateUser(user *entity.UserModel, payload *url.Values) error {
	// Convert birth
	if birth := payload.Get("birth"); birth != "" {
		birthFormatted, err := time.Parse(config.MonthDayYear, birth)
		if err != nil {
			return err
		}
		user.Birth = birthFormatted
	}

	// Hash password
	if password := payload.Get("password"); password != "" {
		hash64, err := helper.HashString(password)
		if err != nil {
			return err
		}

		user.Password = hash64
	}

	// Set role
	if payloadRole := payload.Get("role"); payloadRole != "" {
		isValid := false

		for _, role := range entity.UserRoles {
			if payloadRole == role {
				isValid = true
				break
			}
		}

		if !isValid {
			return errors.New("User role is not valid")
		}
	}

	user.UpdatedAt = time.Now()

	err := user.Validate()
	if err != nil {
		return err
	}

	return nil
}
