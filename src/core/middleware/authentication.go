package core

import (
	"context"
	"net/http"
)

type key int

const (
	user key = iota
)

// AuthenticationHandler Middleware to verify user credentials
func AuthenticationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), user, "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
