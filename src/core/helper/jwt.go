package core

import (
	"os"
	"strconv"
	"time"

	entity "github.com/damascus-mx/photon-api/src/entity"
	jwt "github.com/dgrijalva/jwt-go"
)

// JWTSecret JWT Secret key
var JWTSecret []byte = []byte(os.Getenv("JWT_SECRET"))

// Claims JWT Claims for user
type Claims struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Birth     time.Time `json:"birth"`
	Username  string    `json:"username"`
	Image     *string   `json:"image"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	jwt.StandardClaims
}

// GenerateJWT Creates a new JWT token using Base64 encoding
func GenerateJWT(user *entity.UserModel) (string, error) {
	claims := &Claims{
		ID:        user.ID,
		Name:      user.Name,
		Surname:   user.Surname,
		Birth:     user.Birth,
		Username:  user.Username,
		Image:     user.Image,
		Role:      user.Role,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(0, 0, 14).Unix(),
			Issuer:    "damascus-engineering.com",
			NotBefore: time.Now().Unix(),
			Subject:   strconv.FormatInt(user.ID, 36),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(JWTSecret)
}

// ParseUser Parses JWT claims to a user
func ParseUser(claims *Claims) *entity.UserModel {
	user := new(entity.UserModel)
	user.ID = claims.ID
	user.Name = claims.Name
	user.Birth = claims.Birth
	user.Username = claims.Username
	user.Password = ""
	user.Image = claims.Image
	user.Role = claims.Role
	user.Active = claims.Active
	user.CreatedAt = claims.CreatedAt
	user.UpdatedAt = claims.UpdatedAt

	return user
}
