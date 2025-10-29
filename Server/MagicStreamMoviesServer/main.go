package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// we use the github link because it is best practice
	routes "github.com/VasosTheBeast/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
)

func main() {
	fmt.Println("Niaou go, Niaou GO")

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello MagicStreamMovies!")
	})

	// set up Unprotected and Protected routes
	routes.SetUpUnProtectedRoutes(router)
	routes.SetUpProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
