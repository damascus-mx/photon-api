package entity

import (
	"errors"
	"strconv"
	"time"
)

// MonthDayYear default time pattern MM/DD/YYYY
const monthDayYear string = "01/02/2006"

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

// UserPayload User for payloads, all fields are strings
type UserPayload struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Birth     string `json:"birth"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Image     string `json:"image"`
	Role      string `json:"role"`
	Active    string `json:"active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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

// NewUser Returns an instance of a user model
func NewUser() *UserModel {
	return &UserModel{}
}

// -- LOGICAL SECTION --

// Validate Verify model
func (u *UserModel) Validate() error {
	if u.Name == "" || u.Surname == "" || u.Username == "" || u.Password == "" ||
		!u.Birth.IsZero() {
		return nil
	}

	return errors.New("User is not valid")
}

// Validate Verify model
func (u *UserPayload) Validate() error {
	if u.Name == "" && u.Surname == "" && u.Birth == "" && u.Username == "" && u.Password == "" &&
		u.Image == "" && u.Role == "" && u.Active == "" {
		return errors.New("User Payload is not valid")
	}

	return nil
}

// SetBirth Sets given birth string to the user's memory address
func (u *UserModel) SetBirth(birthPayload string) error {
	birthFormatted, err := time.Parse(monthDayYear, birthPayload)
	if err != nil {
		return err
	}

	u.Birth = birthFormatted
	return nil
}

// SetRole Sets the given role to the user's memory address
func (u *UserModel) SetRole(rolePayload string) error {
	isValid := false

	for _, role := range UserRoles {
		if rolePayload == role {
			isValid = true
			break
		}
	}

	if !isValid {
		return errors.New("User role is not valid")
	}

	u.Role = rolePayload
	return nil
}

// SetActive Sets the given user status to the user's memory address
func (u *UserModel) SetActive(activePayload string) error {
	isActive, err := strconv.ParseBool(activePayload)
	if err != nil {
		return errors.New("User status is not valid")
	}

	u.Active = isActive
	return nil
}
