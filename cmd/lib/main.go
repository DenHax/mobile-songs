package main

import (
	"net/http"


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
