package entity

import "time"

// UserModel User Model
type UserModel struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Image     string    `json:"image"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	Birth     time.Time `json:"birth"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate Verify model
func (user *UserModel) Validate() error {
	return nil
}

// NewUser Returns an instance of a usermodel
func NewUser() *UserModel {
	return &UserModel{}
}
