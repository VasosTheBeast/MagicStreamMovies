package routes

import (
	"github.com/gin-gonic/gin"
	// we use the github link because it is best practice
	controller "github.com/VasosTheBeast/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
)

func SetUpUnProtectedRoutes(router *gin.Engine) {
	// router to get movies - utlizes the controller.getMovies function,
	router.GET("/movies", controller.GetMovies())
	// router to add a user
	router.POST("/register", controller.RegisterUser())
	// router to login
	router.POST("/login", controller.LoginUser())
}
