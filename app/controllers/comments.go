package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"capture-life-api/app/models"
)

// PostCreateComment is for creating comment
func PostCreateComment(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	comment := models.Comment{}
	comment.Title = req.FormValue("title")
	comment.Content = req.FormValue("content")
	comment.Author = req.FormValue("author")
	comment.Timestamp = req.FormValue("timestamp")
	comment.BlogID = req.FormValue("blog")
	comment.ReplyTo = req.FormValue("replyto")
	_, err := db.Collection("comments").InsertOne(context.TODO(), comment)
	if err != nil {
		log.Fatal(err)
	}
	respondJSON(res, http.StatusOK, comment)
}

// GetComments is for getting all comments
func GetComments(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var comments []*models.Comment
	cur, err := db.Collection("comments").Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Comment
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	respondJSON(res, http.StatusOK, comments)
}

// GetComment is for getting a single comment
func GetComment(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	comment := models.Comment{}
	objID, err := primitive.ObjectIDFromHex(params["_id"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	decoder := db.Collection("comments").FindOne(context.TODO(), filter)
	decoder.Decode(&comment)
	respondJSON(res, http.StatusOK, comment)
}

// GetCommentByBlog is for getting a comments for a single blog
func GetCommentByBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	blogID, err := primitive.ObjectIDFromHex(params["blog"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"blog": blogID}

	var comments []*models.Comment
	cur, err := db.Collection("comments").Find(context.TODO(), filter, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Comment
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	respondJSON(res, http.StatusOK, comments)
}

// DeleteComment is for deleting a single comment
func DeleteComment(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	objID, err := primitive.ObjectIDFromHex(params["_id"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	_, err = db.Collection("comments").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	deleteResult := models.Response{}
	deleteResult.Message = "Comment deleted successfully"
	respondJSON(res, http.StatusOK, deleteResult)
}

// DeleteCommentsByBlog is for deleting all comments for a single blog
func DeleteCommentsByBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	objID, err := primitive.ObjectIDFromHex(params["blog"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"blog": objID}
	_, err = db.Collection("comments").DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	deleteResult := models.Response{}
	deleteResult.Message = "Comment deleted successfully"
	respondJSON(res, http.StatusOK, deleteResult)
}

// PutUpdateComment is for updating a single comment
func PutUpdateComment(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	comment := models.Comment{}
	comment.Title = req.FormValue("title")
	comment.Content = req.FormValue("content")
	comment.Author = req.FormValue("author")
	comment.Timestamp = req.FormValue("timestamp")
	comment.BlogID = req.FormValue("blog")
	comment.ReplyTo = req.FormValue("replyto")
	objID, err := primitive.ObjectIDFromHex(params["_id"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"title":     req.FormValue("title"),
		"content":   req.FormValue("content"),
		"author":    req.FormValue("author"),
		"timestamp": req.FormValue("timestamp"),
		//"blog":      req.FormValue("blog"),
		//"replyto":   req.FormValue("replyto"),
	},
	}
	db.Collection("comments").FindOneAndUpdate(context.TODO(), filter, update)
	updateResponse := models.Response{}
	updateResponse.Message = "Comment updated successfully!"
	respondJSON(res, http.StatusOK, updateResponse)
}
