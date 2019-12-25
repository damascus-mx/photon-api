package core

import "github.com/go-chi/chi"

type IRoute interface {
	SetRoutes(router *chi.Mux)
}
