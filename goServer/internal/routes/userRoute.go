package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"goServer/internal/controllers"
	"goServer/internal/middleware"
)

func RegisterUserRoutes(router *mux.Router) {

	// Public routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/refresh", controllers.Refresh).Methods("POST")

	// Protected routes
	router.Handle("/logout",
		middleware.Auth(http.HandlerFunc(controllers.Logout)),
	).Methods("POST", "OPTIONS")

	router.Handle("/me",
		middleware.Auth(http.HandlerFunc(controllers.GetCurrentUser)),
	).Methods("GET", "OPTIONS")
}
