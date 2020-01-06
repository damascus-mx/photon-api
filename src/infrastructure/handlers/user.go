package infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	md "github.com/damascus-mx/photon-api/src/core/middleware"
	utils "github.com/damascus-mx/photon-api/src/core/util"
	"github.com/damascus-mx/photon-api/src/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// UserUsecase User usecase interface
type UserUsecase interface {
	CreateUser(name, surname, username, password, birth string) (int, error)
	GetAllUsers(limit, index int64) ([]*entity.UserModel, error)
	GetUserByID(id int64) (*entity.UserModel, error)
	DeleteUser(id int64) error
	UpdateUser(user *entity.UserModel, payload *url.Values) error
}

// UserHandler HTTP Handler for user
type UserHandler struct {
	userUsecase UserUsecase
}

type key int

const (
	userCtx key = iota
)

// NewUserHandler Returns a handler instance
func NewUserHandler(userCase UserUsecase) *UserHandler {
	return &UserHandler{userCase}
}

// Routes Exports all routes
func (u *UserHandler) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.With(md.PaginateHandler).Get("/", u.getAll)
	router.Post("/", u.create)

	router.Route("/{userID}", func(r chi.Router) {
		r.With(u.userContext).Get("/", u.getByID)
		r.Delete("/", u.delete)
		r.With(u.userContext).Put("/", u.delete)
	})

	return router
}

// ---- USER OPERATIONS ----

// Create user
func (u *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	// Save user
	userID, err := u.userUsecase.CreateUser(r.FormValue("name"), r.FormValue("surname"), r.FormValue("username"),
		r.FormValue("password"), r.FormValue("birth"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: err.Error()})
		return
	}

	render.JSON(w, r, &utils.ResponseModel{Message: fmt.Sprintf("User %d successfully created", userID)})
}

// Get all users
func (u *UserHandler) getAll(w http.ResponseWriter, r *http.Request) {
	params, ok := r.Context().Value("paginateParams").(*md.PaginateParams)
	if !ok {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(500)})
		return
	}

	users, err := u.userUsecase.GetAllUsers(params.Limit, params.Index)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &utils.ResponseModel{Message: err.Error()})
		return
	} else if len(users) <= 0 {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &utils.ResponseModel{Message: "Users not found"})
		return
	}

	render.JSON(w, r, users)
}

// Get user by ID or username
func (u *UserHandler) getByID(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*entity.UserModel)
	if !ok {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(500)})
		return
	}

	render.JSON(w, r, user)
}

// Delete user by ID
func (u *UserHandler) delete(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: "Invalid user ID"})
		return
	}

	err = u.userUsecase.DeleteUser(userID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &utils.ResponseModel{Message: err.Error()})
		return
	}

	render.JSON(w, r, &utils.ResponseModel{Message: fmt.Sprintf("User %d successfully deleted", userID)})
}

func (u *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*entity.UserModel)
	if !ok {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(500)})
		return
	}

	err := u.userUsecase.UpdateUser(user, &r.PostForm)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: err.Error()})
		return
	}

	render.JSON(w, r, &utils.ResponseModel{Message: fmt.Sprintf("User %d successfully updated", user.ID)})
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

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
