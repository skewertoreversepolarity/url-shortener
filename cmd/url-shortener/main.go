package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/skewertoreversepolarity/url-shortener.git/cmd/internal/config"
	"github.com/skewertoreversepolarity/url-shortener.git/cmd/internal/lib/logger/sl"
	"github.com/skewertoreversepolarity/url-shortener.git/cmd/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatal("Error loading local.env file")
	}
	fmt.Println("Environment variables loaded from local.env")

	cfg := config.MustLoad()

	fmt.Println(cfg)
	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener service...", slog.String("env", cfg.Env))

	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StPath)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	//TODO: init router: chi, "chi render"

	//TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
