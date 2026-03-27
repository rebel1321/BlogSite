package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"goServer/config"
	"goServer/internal/models"
	"goServer/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// ✅ FIX: Use function instead of global variable
func getUserCollection() *mongo.Collection {
	if config.DB == nil {
		log.Fatal("❌ Database not initialized")
	}
	return config.DB.Collection("users")
}

// ================= REGISTER =================
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	collection := getUserCollection()

	var existing models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&existing)
	if err == nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().Unix(),
	}

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

// ================= LOGIN =================
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	collection := getUserCollection()

	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.Email)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"email": user.Email},
		bson.M{"$set": bson.M{"refreshToken": refreshToken}},
	)
	if err != nil {
		http.Error(w, "Failed to store refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

// ================= REFRESH =================
func Refresh(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := jwt.Parse(req.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte("refresh_secret"), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	email, ok := claims["email"].(string)
	if !ok {
		http.Error(w, "Invalid token data", http.StatusUnauthorized)
		return
	}

	collection := getUserCollection()

	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{
		"email":        email,
		"refreshToken": req.RefreshToken,
	}).Decode(&user)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken, _, err := utils.GenerateTokens(email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
	})
}

// ================= LOGOUT =================
func Logout(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value("email").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	collection := getUserCollection()

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"email": email},
		bson.M{"$unset": bson.M{"refreshToken": ""}},
	)

	if err != nil {
		http.Error(w, "Logout failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}

// ================= CURRENT USER =================
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value("email").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	collection := getUserCollection()

	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Password = ""
	user.RefreshToken = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
