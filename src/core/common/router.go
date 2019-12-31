package core

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	config "github.com/damascus-mx/photon-api/src/core/config"
	contracts "github.com/damascus-mx/photon-api/src/core/interfaces"
	usecase "github.com/damascus-mx/photon-api/src/usecases"
)

// IRouter HTTP Router interface
type IRouter interface {
	InitializeRouter(router *chi.Mux) *chi.Mux
	SetRoutes(router *chi.Mux)
}

// Router HTTP Router implementation
type Router struct{}

// InitializeRouter Starts the given router with needed configs
func (r *Router) InitializeRouter(router *chi.Mux) *chi.Mux {
	// Get CORS policies
	cors := config.GetCORS()

	router.Use(
		cors.Handler, //	Use user-defined CORS policies

		middleware.RequestID,               //	Set an ID to every request
		middleware.Logger,                  //	Log API request
		middleware.RealIP,                  //	Gets the real IP Address from request
		middleware.Recoverer,               //	Recover from panics without crashing server
		middleware.DefaultCompress,         //	Compress results, gzip assets and json
		middleware.RedirectSlashes,         //	Redirect slashes to no slash URL versions
		middleware.Timeout(60*time.Second), //	Set a timeout value on the request context (ctx)

		render.SetContentType(render.ContentTypeJSON), //	Set HTTP response's headers to application/json
	)

	return router
}

// SetRoutes Mounts resources into the given router
func (r *Router) SetRoutes(router *chi.Mux) {
	var user contracts.Usecase = &usecase.UserUsecase{}
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/user", user.Routes())
	})
}
