package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	migrations "github.com/DenHax/mobile-songs/cmd/migrate"
	"github.com/DenHax/mobile-songs/internal/api/handler"
	"github.com/DenHax/mobile-songs/internal/config"
	"github.com/DenHax/mobile-songs/internal/repo"
	"github.com/DenHax/mobile-songs/internal/server"
	"github.com/DenHax/mobile-songs/internal/service"
	"github.com/DenHax/mobile-songs/internal/storage"
)

type Song struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

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

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Logger.Env)
	log.Info(
		"start logger in",
		slog.String("env", cfg.Logger.Env),
	)

	migrations.Migrate(cfg.Storage.URL, cfg.Storage.MigrationPath, log)

	storage, err := storage.New(cfg.Storage.URL)
	if err != nil {
		log.Error("failed to init storage", slog.String("error", err.Error()))
		os.Exit(1)
	}

	repos := repo.NewRepository(storage)

	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	// TODO: auth
	// TODO: auth-session
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("starting server", slog.String("address", cfg.Server.Address))

	srv := server.New(cfg.Server, handlers.Init())

	go func() {
		if err := srv.Run(); err != nil {
			log.Error("failed to stop server", slog.String("error", err.Error()))
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", slog.String("error", err.Error()))
		return
	}
}
