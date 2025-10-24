package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// we use the github link because it is best practice
	controller "github.com/VasosTheBeast/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
)

func main() {
	fmt.Println("Niaou go, Niaou GO")

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello MagicStreamMovies!")
	})

	// router to get movies - utlizes the controller.getMovies function,
	router.GET("/movies", controller.GetMovies())
	// router to get movie - takes a parameter: the imdb_id, utilizes the getMovie function
	router.GET("/movie/:imdb_id", controller.GetMovie())
	// router to add a movie
	router.POST("/addmovie", controller.AddMovie())
	// router to add a user
	router.POST("/register", controller.RegisterUser())
	// router to login
	router.POST("/login", controller.LoginUser())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
