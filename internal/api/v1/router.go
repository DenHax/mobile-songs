package v1

import (
	"github.com/DenHax/mobile-songs/internal/api/v1/handlers"
	"github.com/gorilla/mux"
)

func NewApi(r *mux.Router) *mux.Router {
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/", handlers.GetPath)
	songs := v1.PathPrefix("/songs").Subrouter()
	songs.HandleFunc("/", handlers.CreateSong).Methods("POST")

	songs.HandleFunc("/{id}", handlers.GetSongs).Methods("GET")
	songs.HandleFunc("/list", handlers.GetSongs).Methods("GET")
	return v1
}
