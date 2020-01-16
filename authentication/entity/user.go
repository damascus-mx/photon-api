package entity

import (
	"errors"
	"time"
)

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

func (u *UserModel) Validate() error {
	if u.Password == "" {
		return errors.New("User not valid")
	}

	return nil
}
