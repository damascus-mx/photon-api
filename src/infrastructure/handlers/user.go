package infrastructure

import (
	"context"
	"net/http"
	"strconv"

	utils "github.com/damascus-mx/photon-api/src/core/util"
	"github.com/damascus-mx/photon-api/src/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// UserUsecase User usecase interface
type UserUsecase interface {
	CreateUser(name, surname, username, password, birth string) (int, error)
	GetAllUsers() ([]*entity.UserModel, error)
	GetUserByID(id int64) (*entity.UserModel, error)
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

	router.Route("/{userID}", func(r chi.Router) {
		r.Use(u.userContext)
		r.Get("/", u.getByID)
	})

	return router
}

// ---- USER OPERATIONS ----

// Create user
func (u *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	// Save user
	_, err := u.userUsecase.CreateUser(r.FormValue("name"), r.FormValue("surname"), r.FormValue("username"),
		r.FormValue("password"), r.FormValue("birth"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: err.Error()})
		return
	}

	render.JSON(w, r, utils.ResponseModel{Message: "User created"})
}

// Get all users
func (u *UserHandler) getAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.userUsecase.GetAllUsers()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &utils.ResponseModel{Message: err.Error()})
		return
	}

	render.JSON(w, r, users)
}

// Get user by ID or username
func (u *UserHandler) getByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value("user").(*entity.UserModel)
	if !ok {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(500)})
		return
	}

	render.JSON(w, r, user)
}

// --> USER CONTEXT <--
// Handler as middelware to retrieve user with ID, sets user into context
func (u *UserHandler) userContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

		user, err := u.userUsecase.GetUserByID(userID)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, &utils.ResponseModel{Message: "User not found"})
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
