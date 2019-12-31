package usecases

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"

	models "github.com/damascus-mx/photon-api/src/infrastructure/models"
)

// UserUsecase User usecase
type UserUsecase struct {
}

// Routes Exports all routes
func (userCase *UserUsecase) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", userCase.getAll)
	return router
}

func (userCase *UserUsecase) getAll(w http.ResponseWriter, r *http.Request) {
	users := []models.UserModel{
		{
			ID:       1,
			Username: "aruizmx",
		},
	}
	render.JSON(w, r, users)
}
