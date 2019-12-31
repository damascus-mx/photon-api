package core

import "github.com/go-chi/chi"

// Usecase Usecase interface
type Usecase interface {
	Routes() *chi.Mux
}
