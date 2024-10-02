package handler

import (
	"encoding/json"
	"net/http"

	v1 "github.com/DenHax/mobile-songs/internal/api/v1"
	"github.com/DenHax/mobile-songs/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

type Response struct {
	Name string `json:"name"`
}

func GetPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Name: "u name"}
	json.NewEncoder(w).Encode(response)
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) Init() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	// api.Use(middleware.HeaderMiddleware)
	api = v1.NewApi(api)
	return r
}
