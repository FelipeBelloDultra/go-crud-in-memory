package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/FelipeBelloDultra/go-crud-in-memory/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write json data", "error", err)
		return
	}
}

func NewHandler(db database.Application) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Post("/users", handleCreateUser(db))
		r.Get("/users", handleListUsers(db))
		r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {})
		r.Put("/users/{id}", func(w http.ResponseWriter, r *http.Request) {})
		r.Delete("/users/{id}", func(w http.ResponseWriter, r *http.Request) {})
	})

	return r
}

type CreateUserRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

func handleCreateUser(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body CreateUserRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(
				w,
				Response{Error: "Please provide FirstName LastName and bio for the user"},
				http.StatusBadRequest,
			)
			return
		}

		if len(body.FirstName) < 2 || len(body.FirstName) > 50 {
			sendJSON(
				w,
				Response{Error: "First name must be between 2 and 50 characters long"},
				http.StatusBadRequest,
			)
			return
		}
		if len(body.LastName) < 2 || len(body.LastName) > 50 {
			sendJSON(
				w,
				Response{Error: "Last name must be between 2 and 50 characters long"},
				http.StatusBadRequest,
			)
			return
		}
		if len(body.Biography) < 20 || len(body.Biography) > 450 {
			sendJSON(
				w,
				Response{Error: "Biography must be between 20 and 450 characters long"},
				http.StatusBadRequest,
			)
			return
		}

		id := db.CreateUser(body.FirstName, body.LastName, body.Biography)

		sendJSON(w, Response{
			Data: id,
		}, http.StatusCreated)
	}
}

func handleListUsers(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userResponse []UserResponse
		users := db.ListUsers()

		for id, u := range users {
			userResponse = append(userResponse, UserResponse{
				ID:        id,
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Biography: u.Biography,
			})
		}

		sendJSON(
			w,
			Response{
				Data: userResponse,
			},
			http.StatusOK,
		)
	}
}
