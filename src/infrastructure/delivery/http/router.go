package infrastructure

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-redis/redis/v7"

	config "github.com/damascus-mx/photon-api/src/core/config"
	handler "github.com/damascus-mx/photon-api/src/infrastructure/handlers"
	repository "github.com/damascus-mx/photon-api/src/infrastructure/repositories"
	usecase "github.com/damascus-mx/photon-api/src/usecases"
)

// IRouter HTTP Router interface
type IRouter interface {
	InitializeRouter(router *chi.Mux) *chi.Mux
	SetRoutes(router *chi.Mux)
}

// Router HTTP Router implementation
type Router struct {
	DB    *sql.DB
	Redis *redis.Client
}

// NewHTTPRouter Returns an HTTP Router instance
func NewHTTPRouter(db *sql.DB, redis *redis.Client) *Router {
	return &Router{db, redis}
}

// InitializeRouter Starts the given router with required configs
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
	router.Route("/v1", func(routerChi chi.Router) {
		routerChi.Mount("/user", handler.NewUserHandler(usecase.NewUserUsecase(repository.NewUserRepository(r.DB))).Routes())
	})
}
