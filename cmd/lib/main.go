package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/DenHax/mobile-songs/internal/config"
	"github.com/DenHax/mobile-songs/internal/storage"
)

type Song struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

var (
	songs = make(map[string]Song)
	mu    sync.Mutex
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func getSongs(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var songList []Song
	for _, song := range songs {
		songList = append(songList, song)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songList)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/songs/"):]

	mu.Lock()
	song, exists := songs[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

func createSong(w http.ResponseWriter, r *http.Request) {
	var song Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	songs[song.Song] = song
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}
func router(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/songs":
		http.HandleFunc("/songs", createSong)
	case "/songs/":
		http.HandleFunc("/songs/", getSong)
	case "/songs/list":
		http.HandleFunc("/songs/list", getSongs)
	default:
		http.NotFound(w, r)
	}
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

func main() {

	// TODO: config
	cfg := config.MustLoad()
	// TODO: logger
	log := setupLogger(cfg.Logger.Env)
	log.Info(
		"start logger in",
		slog.String("env", cfg.Logger.Env),
	)

	// TODO: storage
	storage, err := storage.New(cfg.Storage.URL)
	if err != nil {
		log.Error("failed to init storage", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// TODO: auth
	// TODO: auth-session
	// TODO: repository
	// TODO: server (api | handler)
	// TODO: server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("starting server", slog.String("address", cfg.Server.Address))

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      http.HandlerFunc(router),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
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
