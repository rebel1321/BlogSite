package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"goServer/internal/controllers"
	"goServer/internal/middleware"
)

func RegisterPostRoutes(router *mux.Router) {

	// Create (protected)
	router.Handle("/posts",
		middleware.Auth(http.HandlerFunc(controllers.CreatePost)),
	).Methods("POST")

	// Read
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{slug}", controllers.GetPost).Methods("GET")

	// Update (protected)
	router.Handle("/posts/{slug}",
		middleware.Auth(http.HandlerFunc(controllers.UpdatePost)),
	).Methods("PUT")

	// Delete (protected)
	router.Handle("/posts/{slug}",
		middleware.Auth(http.HandlerFunc(controllers.DeletePost)),
	).Methods("DELETE")
}