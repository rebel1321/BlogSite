package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"goServer/config"
	"goServer/internal/cron"
	"goServer/internal/routes"
)

func Start() {
	// Load environment variables (optional - in production, env vars are set by platform)
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println("⚠️  No .env file found - using platform environment variables")
	}

	// Connect to database
	config.ConnectDB()

	// Connect to cloudinary
	config.InitCloudinary()

	// Start cron jobs
	cron.StartCronJobs()
	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Setup routes
	router := routes.SetupRoutes()

	fmt.Println("🚀 Server running on port", port)

	// Start server
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}