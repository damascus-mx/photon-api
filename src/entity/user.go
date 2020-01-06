package entity

import (
	"errors"
	"time"
)

// UserModel User Model
type UserModel struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Birth     time.Time `json:"birth"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Image     *string   `json:"image"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRoles User roles enum
var UserRoles = [...]string{
	"ROLE_STUDENT",
	"ROLE_TEACHER",
	"ROLE_OPERATOR",
	"ROLE_MANAGER",
	"ROLE_ADMIN",
	"ROLE_ROOT",
	"ROLE_SUPPORT",
	"ROLE_EMPLOYEE",
	"ROLE_PRINCIPAL",
}

// Validate Verify model
func (u *UserModel) Validate() error {
	if u.Name == "" || u.Surname == "" || u.Username == "" || u.Password == "" ||
		!u.Birth.IsZero() {
		return nil
	}

	return errors.New("User is not valid")
}

// NewUser Returns an instance of a user model
func NewUser() *UserModel {
	return &UserModel{}
}
