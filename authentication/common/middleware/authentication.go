package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"

	helper "github.com/damascus-mx/photon-api/authentication/common/helper"
	util "github.com/damascus-mx/photon-api/authentication/common/util"
)

// AuthenticationHandler Middleware to verify user credentials
func AuthenticationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if bearerString := strings.Split(r.Header.Get("Authorization"), " ")[0]; bearerString != "Bearer" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, &util.ResponseModel{Message: http.StatusText(http.StatusUnauthorized)})
			return
		}

		claims := new(helper.Claims)

		tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
		token, err := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
			return helper.JWTSecret, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, &util.ResponseModel{Message: http.StatusText(http.StatusUnauthorized)})
				return
			}

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, &util.ResponseModel{Message: "Invalid Token"})
			return
		}

		if !token.Valid {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, &util.ResponseModel{Message: http.StatusText(http.StatusUnauthorized)})
			return
		}

		ctx := context.WithValue(r.Context(), "user", helper.ParseUser(claims))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
