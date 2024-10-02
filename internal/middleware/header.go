package middleware

import "net/http"

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Server", "LibreLib")
		w.Header().Set("Content-Type", "application/json")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		next.ServeHTTP(w, r)
	})
}
