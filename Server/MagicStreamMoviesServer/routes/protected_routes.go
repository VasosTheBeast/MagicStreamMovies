package routes

import (
	"github.com/gin-gonic/gin"
	// we use the github link because it is best practice
	controller "github.com/VasosTheBeast/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/VasosTheBeast/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
)

// function that sets up protected routes
// gets as input a router which is a *gin.Engine
func SetUpProtectedRoutes(router *gin.Engine) {
	// force the router to use this function - this is like an intermediate or Middleware endpoint
	// where we check the token (if the user is authorized)
	router.Use(middleware.AuthMiddleware())
	// router to get movie - takes a parameter: the imdb_id, utilizes the getMovie function
	router.GET("/movie/:imdb_id", controller.GetMovie())
	// router to add a movie
	router.POST("/addmovie", controller.AddMovie())

}
