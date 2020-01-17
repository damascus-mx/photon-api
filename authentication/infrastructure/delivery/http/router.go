package delivery

import (
	"database/sql"

	handler "github.com/damascus-mx/photon-api/authentication/infrastructure/handler/http"
	"github.com/damascus-mx/photon-api/authentication/infrastructure/repository"
	"github.com/damascus-mx/photon-api/authentication/usecase"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v7"
)

// HTTPRouter HTTP Router
type HTTPRouter struct {
	DB    *sql.DB
	Cache *redis.Client
}

// NewRouter Create new HTTP Router
func (r *HTTPRouter) NewRouter(database *sql.DB, redisClient *redis.Client) *chi.Mux {
	r.DB = database
	r.Cache = redisClient

	router := chi.NewRouter()

	router.Route("/v1", func(routerChi chi.Router) {
		// Auth Dependency Injection
		routerChi.Mount("/authentication/user", handler.NewAuthHandler(usecase.NewAuthUsecase(repository.NewAuthRepository(r.Cache), repository.NewUserRepository(r.DB, r.Cache))).Routes())
	})
	return router
}
