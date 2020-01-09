package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

const port = ":8080"

// Task Task Model
type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// Response generic response message
type Response struct {
	Message string `json:"message"`
}

// Server Main logical handler
type Server struct{}

// GetAll Obtain all tasks
func (s *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	var tasks [5]*Task
	for index := 1; index <= 5; index++ {
		task := new(Task)
		task.ID = index
		task.Name = fmt.Sprintf("Lab task %d", index)
		task.Username = "James Bell"

		tasks[index-1] = task
	}

	render.JSON(w, r, tasks)
}

// InitApplication Start router
func (s *Server) InitApplication() *chi.Mux {
	fmt.Print("Running Photon Tasks Microservice\n")
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

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

	router.Get("/task", s.GetAll)
	return router
}

func main() {
	s := new(Server)
	err := http.ListenAndServe(port, s.InitApplication())
	if err != nil {
		log.Fatalf("Error during router init: \n %v", err)
	}

}
