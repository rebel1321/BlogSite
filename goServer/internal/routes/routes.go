package routes

import (
	"net/http"
	"os"

	"goServer/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Global middlewares
	router.Use(middleware.Cors)
	// router.Use(middleware.Logger)

	// Root endpoint
	router.HandleFunc("/", rootHandler).Methods("GET", "OPTIONS")

	// Health check (no prefix)
	router.HandleFunc("/health", healthCheckHandler).Methods("GET", "OPTIONS")

	// API group
	api := router.PathPrefix("/api").Subrouter()

	// Handle OPTIONS for all /api/* routes
	api.HandleFunc("/{path:.*}", handleOptions).Methods("OPTIONS")

	// Register all module routes here
	RegisterUserRoutes(api)
	RegisterPostRoutes(api)

	return router
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "🚀 MegaBlog API Server", "status": "running", "endpoints": {"/health": "Server health check", "/api/users": "User management", "/api/posts": "Blog posts"}}`))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy"}`))
}

// Handle OPTIONS preflight requests for CORS
func handleOptions(w http.ResponseWriter, r *http.Request) {
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}
	w.Header().Set("Access-Control-Allow-Origin", frontendURL)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}
