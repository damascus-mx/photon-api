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
	redis *redis.Client
}

func (r *HTTPRouter) NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(routerChi chi.Router) {
		routerChi.Mount("/authentication", handler.NewAuthHandler(usecase.NewAuthUsecase(repository.NewAuthRepository(r.redis), repository.NewUserRepository(r.DB, r.redis))).Routes())
	})
	return router
}
