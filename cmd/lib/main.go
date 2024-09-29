package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/DenHax/mobile-songs/internal/config"
)

type Song struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

var (
	songs = make(map[string]Song)
	mu    sync.Mutex
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

func main() {

	// TODO: config
	cfg := config.MustLoad()
	// TODO: logger
	// TODO: storage
	// TODO: server


	http.ListenAndServe(":8080", nil)
}
