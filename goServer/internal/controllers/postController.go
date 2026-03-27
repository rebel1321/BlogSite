package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"goServer/config"
	"goServer/internal/models"
	"goServer/internal/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// helper
func getPostCollection() *mongo.Collection {
	if config.DB == nil {
		log.Fatal("DB not initialized")
	}
	return config.DB.Collection("posts")
}

// ================= CREATE POST =================
func CreatePost(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	title := r.FormValue("title")
	slug := r.FormValue("slug")
	content := r.FormValue("content")
	status := r.FormValue("status")

	if title == "" || slug == "" {
		http.Error(w, "Title and slug required", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("email").(string)

	imageURL, imageID, err := services.UploadFile(file)
	if err != nil {
		http.Error(w, "Image upload failed", http.StatusInternalServerError)
		return
	}

	post := models.Post{
		Title:     title,
		Slug:      slug,
		Content:   content,
		ImageURL:  imageURL,
		ImageID:   imageID,
		Status:    status,
		UserID:    userID,
		CreatedAt: time.Now().Unix(),
	}

	_, err = getPostCollection().InsertOne(context.TODO(), post)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// ================= GET SINGLE =================
func GetPost(w http.ResponseWriter, r *http.Request) {

	slug := mux.Vars(r)["slug"]

	var post models.Post
	err := getPostCollection().FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&post)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// ================= GET ALL =================
func GetPosts(w http.ResponseWriter, r *http.Request) {

	cursor, err := getPostCollection().Find(context.TODO(), bson.M{"status": "active"})
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	var posts []models.Post
	if err := cursor.All(context.TODO(), &posts); err != nil {
		http.Error(w, "Error reading posts", http.StatusInternalServerError)
		return
	}

	// Return empty array instead of null if no posts
	if posts == nil {
		posts = []models.Post{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// ================= UPDATE =================
func UpdatePost(w http.ResponseWriter, r *http.Request) {

	slug := mux.Vars(r)["slug"]

	var existing models.Post
	err := getPostCollection().FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&existing)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	updateData := bson.M{}

	title := r.FormValue("title")
	content := r.FormValue("content")
	status := r.FormValue("status")

	if title != "" {
		updateData["title"] = title
	}
	if content != "" {
		updateData["content"] = content
	}
	if status != "" {
		updateData["status"] = status
	}

	// Check new image
	file, _, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		// delete old image
		services.DeleteFile(existing.ImageID)

		imageURL, imageID, err := services.UploadFile(file)
		if err == nil {
			updateData["imageUrl"] = imageURL
			updateData["imageId"] = imageID
		}
	}

	_, err = getPostCollection().UpdateOne(
		context.TODO(),
		bson.M{"slug": slug},
		bson.M{"$set": updateData},
	)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}

	// Fetch the updated post
	var updatedPost models.Post
	err = getPostCollection().FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Failed to fetch updated post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)
}

// ================= DELETE =================
func DeletePost(w http.ResponseWriter, r *http.Request) {

	slug := mux.Vars(r)["slug"]

	var post models.Post
	err := getPostCollection().FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&post)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// delete image
	services.DeleteFile(post.ImageID)

	_, err = getPostCollection().DeleteOne(context.TODO(), bson.M{"slug": slug})
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Post deleted",
	})
}
