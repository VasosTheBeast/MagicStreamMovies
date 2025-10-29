package middleware

import (
	"net/http"

	"github.com/VasosTheBeast/MagicStreamMovies/Server/MagicStreamMoviesServer/utils"
	"github.com/gin-gonic/gin"
)

// gin gonic handler function - validate the access tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the access token
		token, err := utils.GetAccessToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort() // abort the context execution from moving to the next endpoint
			return
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}
		c.Set("userId", claims.UserID)
		c.Set("role", claims.Role)
		c.Next() // continue to the next endpoint
	}
}
