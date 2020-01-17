package usecase

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/damascus-mx/photon-api/authentication/common/helper"
	"github.com/damascus-mx/photon-api/authentication/entity"
	"github.com/dgrijalva/jwt-go"
)

type authRepository interface {
	GetUserToken(username string) (string, error)
	SetUserToken(username, token string) error
}

type userRepository interface {
	FetchByUsername(username string) (*entity.UserModel, error)
}

// AuthUsecase Authentication use cases
type AuthUsecase struct {
	authRepository authRepository
	userRepository userRepository
}

// NewAuthUsecase Create Authentication use case
func NewAuthUsecase(aRepository authRepository, usRepository userRepository) *AuthUsecase {
	return &AuthUsecase{authRepository: aRepository, userRepository: usRepository}
}

// VerifyBearer Return if a Bearer is valid or not
func (a *AuthUsecase) VerifyBearer(bearer string) (int, *entity.UserModel, error) {
	if bearerString := strings.Split(bearer, " ")[0]; bearerString != "Bearer" {
		return http.StatusBadRequest, nil, errors.New("Invalid Bearer Token")
	}

	claims := new(helper.Claims)

	tokenString := strings.Split(bearer, " ")[1]
	token, err := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
		return helper.JWTSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized, nil, err
		}

		return http.StatusBadRequest, nil, errors.New("Invalid Bearer Token")
	}

	if !token.Valid {
		return http.StatusUnauthorized, nil, errors.New("Invalid token")
	}

	return 200, helper.ParseUser(claims), nil
}

// Authenticate Verify a user
func (a *AuthUsecase) Authenticate(username, password string) (string, error) {
	// Get from redis cache
	token, err := a.authRepository.GetUserToken(username)
	if err == nil && token != "" {
		return token, nil
	}

	// Get from DB
	user, err := a.userRepository.FetchByUsername(username)
	if err != nil {
		return "", errors.New("Invalid username/password")
	}

	ok, err := helper.CompareString(password, user.Password)
	if err != nil {
		return "", err
	} else if !ok {
		return "", errors.New("Invalid username/password")
	}

	token, err = helper.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	// Add token to redis
	go func() {
		err = a.authRepository.SetUserToken(username, token)
		if err != nil {
			log.Fatalf("Failed to insert issued token\nerr:%s", err.Error())
		}
	}()

	return token, nil
}
