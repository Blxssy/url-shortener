package main

import (
	"github.com/Blxssy/url-shortener/internal/config"
	"github.com/Blxssy/url-shortener/internal/lib/logger/sl"
	"github.com/Blxssy/url-shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug logging enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("error initializing storage", sl.Err(err))
		os.Exit(1)
	}

	err = storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("error initializing storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// TODO: init router: chi

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
