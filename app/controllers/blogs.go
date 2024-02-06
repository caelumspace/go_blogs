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

// PostCreateBlog is for creating blog
func PostCreateBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	blog := models.Blog{}
	blog.Title = req.FormValue("title")
	blog.Content = req.FormValue("content")
	blog.Author = req.FormValue("author")
	blog.Timestamp = req.FormValue("timestamp")
	_, err := db.Collection("blogs").InsertOne(context.TODO(), blog)
	if err != nil {
		log.Fatal(err)
	}
	respondJSON(res, http.StatusOK, blog)
}

// GetBlogs is for getting all blogs with comments
func GetBlogs(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var blogs []*models.Blog
	cur, err := db.Collection("blogs").Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Blog
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		txt, err := elem.ID.MarshalText()
		filter := bson.M{"blog": string(txt)}
		cur1, err1 := db.Collection("comments").Find(context.TODO(), filter, options.Find())
		if err1 != nil {
			log.Fatal(err1)
		}
		for cur1.Next(context.TODO()) {
			var c models.Comment
			err1 := cur1.Decode(&c)
			if err1 != nil {
				log.Fatal(err1)
			}
			elem.Comments = append(elem.Comments, &c)
		}

		blogs = append(blogs, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	respondJSON(res, http.StatusOK, blogs)
}

// GetBlog is for getting a single blog
func GetBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	blog := models.Blog{}
	objID, err := primitive.ObjectIDFromHex(params["_id"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	decoder := db.Collection("blogs").FindOne(context.TODO(), filter)
	decoder.Decode(&blog)

	txt, err := objID.MarshalText()
	filter1 := bson.M{"blog": string(txt)}
	cur1, err1 := db.Collection("comments").Find(context.TODO(), filter1, options.Find())
	if err1 != nil {
		log.Fatal(err1)
	}
	for cur1.Next(context.TODO()) {
		var c models.Comment
		err1 := cur1.Decode(&c)
		if err1 != nil {
			log.Fatal(err1)
		}
		blog.Comments = append(blog.Comments, &c)
	}

	respondJSON(res, http.StatusOK, blog)
}

// DeleteBlog is for deleting a single blog with comments
func DeleteBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	objID, err := primitive.ObjectIDFromHex(params["_id"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	_, err = db.Collection("blogs").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	deleteResult := models.Response{}
	deleteResult.Message = "Blog deleted successfully"

	// delete comments for blog
	txt, err := objID.MarshalText()
	filter1 := bson.M{"blog": string(txt)}
	_, err = db.Collection("comments").DeleteMany(context.TODO(), filter1)
	if err != nil {
		log.Fatal(err)
	}
	deleteResult = models.Response{}
	deleteResult.Message = "Blog and comments deleted successfully"

	respondJSON(res, http.StatusOK, deleteResult)
}

// PutUpdateBlog is for updating a single blog
func PutUpdateBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	blog := models.Blog{}
	blog.Title = req.FormValue("title")
	blog.Content = req.FormValue("content")
	blog.Author = req.FormValue("author")
	blog.Timestamp = req.FormValue("timestamp")
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
	},
	}
	db.Collection("blogs").FindOneAndUpdate(context.TODO(), filter, update)
	updateResponse := models.Response{}
	updateResponse.Message = "Blog updated successfully!"

	respondJSON(res, http.StatusOK, updateResponse)
}
