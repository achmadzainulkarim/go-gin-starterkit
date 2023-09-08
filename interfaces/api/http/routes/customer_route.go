package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute1Routes(router *gin.Engine) {
	r := router.Group("/route1")
	{
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Route 1"})
		})
		// Add more route handlers as needed.
	}
}
