package main

import (
	"fmt"
	"go-categories-microservice/cmd/api/routes"
	"go-categories-microservice/config"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Time int `json:"time" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/healthy", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": true,
		})
	})

	routes.CategoryRoutes(router)

	router.Run(fmt.Sprintf(":%d", config.API_PORT))
}
