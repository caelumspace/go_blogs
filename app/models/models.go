package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Blog model
type Blog struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	Timestamp string             `json:"timestamp" bson:"timestamp"`
	Comments  []*Comment         `json:"comments" bson:"comments"`
}

// Comment model
type Comment struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	Timestamp string             `json:"timestamp" bson:"timestamp"`
	BlogID    string             `json:"blog" bson:"blog"`
	ReplyTo   string             `json:"replyto" bson:"replyto"`
}

// Response model
type Response struct {
	Message string `json:"message" bson:"message"`
}
