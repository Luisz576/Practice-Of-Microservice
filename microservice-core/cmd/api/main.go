package main

import (
	"fmt"
	"microservice-core/config"

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

	router.Run(fmt.Sprintf(":%d", config.API_PORT))
}
