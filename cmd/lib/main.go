package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Song struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

var (
	songs = make(map[string]Song) // Используем ID песни как ключ
	mu    sync.Mutex
)

func getSongs(w http.ResponseWriter, r *http.Request) {
func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func createSong(w http.ResponseWriter, r *http.Request) {
	}
}

func main() {
	http.HandleFunc("/songs", createSong)
	http.HandleFunc("/songs/", getSong)
	http.HandleFunc("/songs/list", getSongs)

	http.ListenAndServe(":8080", nil)
}
