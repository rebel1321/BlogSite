package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Slug      string             `bson:"slug" json:"slug"`
	Content   string             `bson:"content" json:"content"`
	ImageURL  string             `bson:"imageUrl" json:"imageUrl"`
	ImageID   string             `bson:"imageId" json:"imageId"`
	Status    string             `bson:"status" json:"status"`
	UserID    string             `bson:"userId" json:"userId"`
	CreatedAt int64              `bson:"createdAt" json:"createdAt"`
}