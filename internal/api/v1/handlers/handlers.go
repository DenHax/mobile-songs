package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	mu sync.Mutex
)

func CreateSong(w http.ResponseWriter, r *http.Request) {

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(song)
}

func GetSong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetSongs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type Response struct {
	Name string `json:"name"`
}

func GetPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Name: "api/v1/"}
	json.NewEncoder(w).Encode(response)
	// w.WriteHeader(http.StatusOK)
}
