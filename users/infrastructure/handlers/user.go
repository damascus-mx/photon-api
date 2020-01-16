package infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	md "github.com/damascus-mx/photon-api/users/core/middleware"
	utils "github.com/damascus-mx/photon-api/users/core/util"
	"github.com/damascus-mx/photon-api/users/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// UserUsecase User usecase interface
type UserUsecase interface {
	CreateUser(name, surname, username, password, birth, role string) (int, error)
	GetAllUsers(limit, index int64) ([]*entity.UserModel, error)
	GetUserByID(id int64) (*entity.UserModel, error)
	DeleteUser(id int64) error
	UpdateUser(user *entity.UserModel, payload *entity.UserPayload) error
	AuthenticateUser(username, password string) (string, error)
	GetUserByUsername(username string) (*entity.UserModel, error)
}

// UserHandler HTTP Handler for user
type UserHandler struct {
	userUsecase UserUsecase
}

type (
	key         int
	tknResponse struct {
		Token string `json:"token"`
	}
)

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
	router.Post("/", u.create)
	router.Post("/authorize", u.signIn)
	router.With(md.AuthenticationHandler).With(md.PaginateHandler).Get("/", u.getAll)

	router.Route("/{userID}", func(r chi.Router) {
		r.With(md.AuthenticationHandler).With(u.userContext).Get("/", u.getByID)
		r.With(md.AuthenticationHandler).Delete("/", u.delete)
		r.With(md.AuthenticationHandler).With(u.userContext).Put("/", u.update)
	})

	router.Route("/username/{username}", func(r chi.Router) {
		r.With(md.AuthenticationHandler).Get("/", u.getByUsername)
	})

	return router
}

// ---- USER OPERATIONS ----

// Create user
func (u *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	// Check role
	var role string

	if role = strings.ToUpper(r.FormValue("role")); role != "" {
		switch role {
		case "ROLE_ROOT", "ROLE_SUPPORT":
			if secret := r.URL.Query().Get("secret"); secret == "" || secret != os.Getenv("PHOTON_SECRET") {
				render.Status(r, http.StatusForbidden)
				render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusForbidden)})
				return
			}
		default:
			role = "ROLE_STUDENT"
			break
		}
	} else {
		role = "ROLE_STUDENT"
	}

	// Save user
	userID, err := u.userUsecase.CreateUser(r.FormValue("name"), r.FormValue("surname"), r.FormValue("username"),
		r.FormValue("password"), r.FormValue("birth"), role)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: err.Error()})
		return
	}

	render.JSON(w, r, &utils.ResponseModel{Message: fmt.Sprintf("User %d successfully created", userID)})
}

// Get all users
func (u *UserHandler) getAll(w http.ResponseWriter, r *http.Request) {
	params, ok := r.Context().Value(md.ParamCtx).(*md.PaginateParams)
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

// Update user
func (u *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	loggedUser, ok := r.Context().Value("user").(*entity.UserModel)
	if !ok {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusUnauthorized)})
		return
	}

	user, ok := r.Context().Value(userCtx).(*entity.UserModel)
	if !ok {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(500)})
		return
	}

	// Check role
	switch {
	// Trying to update root user without root account
	case user.Role == "ROLE_ROOT" && loggedUser.Role != "ROLE_ROOT":
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusForbidden)})
		return
	// Trying to update support user without root account
	case loggedUser.ID != user.ID && user.Role == "ROLE_SUPPORT" && loggedUser.Role != "ROLE_ROOT":
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusForbidden)})
		return
	// Trying to update someone else account without root/support account
	case loggedUser.ID != user.ID && user.Role != "ROLE_ROOT" && (loggedUser.Role != "ROLE_ROOT" || loggedUser.Role != "ROLE_SUPPORT"):
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, &utils.ResponseModel{Message: http.StatusText(http.StatusForbidden)})
		return
	}

	payload := new(entity.UserPayload)
	payload.Name = r.FormValue("name")
	payload.Surname = r.FormValue("surname")
	payload.Birth = r.FormValue("birth")
	payload.Username = r.FormValue("username")
	payload.Password = r.FormValue("password")
	payload.Image = r.FormValue("image")
	payload.Role = r.FormValue("role")
	payload.Active = r.FormValue("active")

	err := payload.Validate()
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: "Missing fields to update"})
		return
	}

	err = u.userUsecase.UpdateUser(user, payload)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: err.Error()})
		return
	}

	render.JSON(w, r, &utils.ResponseModel{Message: fmt.Sprintf("User %d successfully updated", user.ID)})
}

// Sign In
func (u *UserHandler) signIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	token, err := u.userUsecase.AuthenticateUser(username, password)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: "Username / Password invalid"})
		return
	}

	render.JSON(w, r, &tknResponse{Token: token})
}

func (u *UserHandler) getByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &utils.ResponseModel{Message: "This resources requires a username"})
		return
	}

	user, err := u.userUsecase.GetUserByUsername(username)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &utils.ResponseModel{Message: "User not found"})
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

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
