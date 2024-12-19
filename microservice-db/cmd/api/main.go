package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/healthy", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": true,
		})
	})

	CategoryRoutes(router)

	router.Run(":8000")
}
