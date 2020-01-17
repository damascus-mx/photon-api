package handler

import (
	"net/http"

	"github.com/damascus-mx/photon-api/authentication/common/util"
	"github.com/damascus-mx/photon-api/authentication/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type authUsecase interface {
	VerifyBearer(bearer string) (int, *entity.UserModel, error)
	Authenticate(username, password string) (string, error)
}

// NewAuthHandler Get a authentication handler
func NewAuthHandler(usecase authUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: usecase}
}

// AuthHandler HTTP Handler for authentication
type AuthHandler struct {
	authUsecase authUsecase
}

// Routes Export all routes
func (a *AuthHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/bearer", a.verifyBearer)
	router.Post("/authorize", a.signIn)
	return router
}

func (a *AuthHandler) verifyBearer(w http.ResponseWriter, r *http.Request) {
	if bearer := r.FormValue("token"); bearer != "" {
		status, user, err := a.authUsecase.VerifyBearer(bearer)
		switch {
		case err == nil && status == 200:
			render.JSON(w, r, user)
			return
		case err != nil && status == http.StatusBadRequest:
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, &util.ResponseModel{Message: err.Error()})
			return
		case err != nil && status == http.StatusUnauthorized:
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, &util.ResponseModel{Message: err.Error()})
			return
		}
	}

	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, &util.ResponseModel{Message: "Bearer Token not found"})
}

func (a *AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" && password == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &util.ResponseModel{Message: "Username and password required"})
		return
	}

	token, err := a.authUsecase.Authenticate(username, password)
	if err != nil || token == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &util.ResponseModel{Message: "Incorrect username/password"})
		return
	}

	render.JSON(w, r, &tokenResponse{Token: token})
}
