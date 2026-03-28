package middleware

import (
	"net/http"
	"os"
	"strings"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get allowed origins from environment variable (comma-separated)
		allowedOriginsStr := os.Getenv("ALLOWED_ORIGINS")
		if allowedOriginsStr == "" {
			allowedOriginsStr = "http://localhost:5173"
		}

		// Parse comma-separated origins
		allowedOrigins := strings.Split(allowedOriginsStr, ",")

		// Get the requesting origin
		origin := r.Header.Get("Origin")

		// Check if origin is allowed
		isAllowed := false
		for _, allowed := range allowedOrigins {
			if strings.TrimSpace(allowed) == origin {
				isAllowed = true
				break
			}
		}

		// Set CORS headers if origin is allowed
		if isAllowed {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
