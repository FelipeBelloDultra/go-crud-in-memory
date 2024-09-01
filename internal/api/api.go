package api

import (
	"net/http"

	"github.com/FelipeBelloDultra/go-crud-in-memory/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandler(db database.Database) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Post("/users", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/users", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {})
		r.Put("/users/{id}", func(w http.ResponseWriter, r *http.Request) {})
		r.Delete("/users/{id}", func(w http.ResponseWriter, r *http.Request) {})
	})

	return r
}
