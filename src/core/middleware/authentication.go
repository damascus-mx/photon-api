package core

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"

	core "github.com/damascus-mx/photon-api/src/core/helper"
	utils "github.com/damascus-mx/photon-api/src/core/util"
)

// AuthenticationHandler Middleware to verify user credentials
func AuthenticationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if bearerString := strings.Split(r.Header.Get("Authorization"), " ")[0]; bearerString != "Bearer" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusUnauthorized)})
			return
		}

		claims := new(core.Claims)

		tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
		token, err := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
			return core.JWTSecret, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusUnauthorized)})
				return
			}

			fmt.Printf("\nError: %s", err.Error())

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusBadRequest)})
			return
		}

		if !token.Valid {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusUnauthorized)})
			return
		}

		ctx := context.WithValue(r.Context(), "user", core.ParseUser(claims))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
