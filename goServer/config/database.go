package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Get MongoDB URI
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}

	// Client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ✅ FIXED LINE
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	fmt.Println("✅ MongoDB Connected")

	DB = client.Database("blogdb")
}