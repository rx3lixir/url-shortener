package main

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rx3lixir/urlshortener/internal/config"
	"github.com/rx3lixir/urlshortener/internal/http-server/handlers/url/save"
	mwLogger "github.com/rx3lixir/urlshortener/internal/http-server/middleware/logger"
	"github.com/rx3lixir/urlshortener/internal/lib/logger/sl"
	"github.com/rx3lixir/urlshortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))

	log.Info("starting server:", "address:", cfg.Address)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.Timeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server crashed")
}

func setupLogger(env string) *log.Logger {
	var logger *log.Logger

	switch env {
	case envLocal:
		logger = log.NewWithOptions(os.Stdout, log.Options{
			Prefix:    "🍃env=local",
			Formatter: log.TextFormatter,
			Level:     log.DebugLevel,
		})
	case envDev:
		logger = log.NewWithOptions(os.Stdout, log.Options{
			Prefix:    "🍃env=dev",
			Formatter: log.JSONFormatter,
			Level:     log.DebugLevel,
		})
	case envProd:
		logger = log.NewWithOptions(os.Stdout, log.Options{
			Prefix:    "🍃env=prod",
			Formatter: log.JSONFormatter,
			Level:     log.InfoLevel,
		})
	}

	return logger
}
