package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/FelipeBelloDultra/go-crud-in-memory/internal/api"
	"github.com/FelipeBelloDultra/go-crud-in-memory/internal/database"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("all systems offline")
}

func run() error {
	database := database.MakeDatabase()
	handler := api.NewHandler(database)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
