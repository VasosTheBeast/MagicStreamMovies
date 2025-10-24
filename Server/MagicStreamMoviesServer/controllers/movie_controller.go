package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/VasosTheBeast/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	"github.com/VasosTheBeast/MagicStreamMovies/Server/MagicStreamMoviesServer/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// load the movies collection from mongo using the database package
var movieCollection *mongo.Collection = database.OpenCollection("movies")

// create a validator object using the validator package we imported, this is used to validate data from the models
var validate = validator.New()

// first letter is capital because we want this function to be importable from other packages
// this function returns another Function: a gin.HanlderFunc:
// gin.HandlerFunc is a function that takes as input a pointer to a gin.Context - which represents the HTTP requests or responses
// and does the logic of the funciton. The returned function is called when a HTTP requests comes in
func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		// create a timeout context to cancel long running queries
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		// context is cleaned up after the funciton finishes
		defer cancel()

		var movies []models.Movie
		// return the query and return a cursor - an iterator over the query results,
		// Find function comes from the mongoDB connection instance. bson.M{} means use an empty filter
		cursor, err := movieCollection.Find(ctx, bson.M{})
		// if there is an error handle it:
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch movies."})
		}
		// make sure cursor will be closed at the end of the function
		defer cursor.Close(ctx)
		// read all documents from the result - using cursor.ALL - and decode them in the movies slice (list)
		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode movies."})
		}
		// Send an HTTP 200 ok response with body the movies serilazied to JSON
		c.JSON(http.StatusOK, movies)
	}
}

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		// context is cleaned up after the funciton finishes
		defer cancel()

		// read the parameter from the http request (c)
		movieID := c.Param("imdb_id")
		if movieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "MovieID is required"})
			return
		}
		var movie models.Movie
		// query the db with the correct filter
		// here we do not have a cursor since we expect only one answer
		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}
		c.JSON(http.StatusOK, movie)
	}
}

func AddMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		// create a timeout Context - when the timeout expires, the ctx is cancelled with any operation listening on it
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		// manually cancel the context when the function ends
		defer cancel()
		// variable to store the movie that will be sent to mongo
		var movie models.Movie
		// c.ShouldBindJSON - attempts to parse the incoming json payload to the movie struct.
		// if there is error (err=nil) -> send a bad request and return
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
			return
		}
		// validate movie with the Movie struct in models using the validator import
		if err := validate.Struct(movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error})
			return
		}

		// insert the movie to mongo using the movieCollection object
		// this inserts a document (a movie) to the mongo DB collection
		// movie will be converted into a bson object and stored in mongoDB
		result, err := movieCollection.InsertOne(ctx, movie)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add movie"})
			return
		}
		// return success message
		c.JSON(http.StatusCreated, result)
	}
}
