package usecase

import (
	entity "github.com/damascus-mx/photon-api/src/entity"
)

// UserRepository Persistence user layer
type UserRepository interface {
	Save(user *entity.UserModel) error
	FetchByID(id int) *entity.UserModel
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

// CreateUser Save a new user
func (u *UserUsecase) CreateUser(user *entity.UserModel) error {
	user.Validate()
	u.userRepository.Save(user)
	return nil
}

// GetAllUsers Retrieves all users
func (u *UserUsecase) GetAllUsers() ([]*entity.UserModel, error) {
	return u.userRepository.FetchAll()
}
