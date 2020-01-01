package infrastructure

import (
	"net/http"

	utils "github.com/damascus-mx/photon-api/src/core/utils"
	"github.com/damascus-mx/photon-api/src/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// UserUsecase User usecase interface
type UserUsecase interface {
	CreateUser(user *entity.UserModel) error
	GetAllUsers() ([]*entity.UserModel, error)
}

// UserHandler HTTP Handler for user
type UserHandler struct {
	userUsecase UserUsecase
}

// NewUserHandler Returns a handler instance
func NewUserHandler(userCase UserUsecase) *UserHandler {
	return &UserHandler{userCase}
}

// Routes Exports all routes
func (u *UserHandler) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", u.getAll)
	router.Post("/", u.create)
	return router
}

func (u *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	user := entity.NewUser()
	user.ID = 1
	user.Username = "aruiz"

	err := u.userUsecase.CreateUser(user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, utils.ResponseModel{Message: err.Error()})
	}

	render.JSON(w, r, utils.ResponseModel{Message: "User created."})
}

func (u *UserHandler) getAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.userUsecase.GetAllUsers()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, utils.ResponseModel{Message: err.Error()})
	}
	render.JSON(w, r, users)
}
