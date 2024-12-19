package main

import (
	"microservice-core/cmd/api/controller"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/", controller.CreateCategory)
	categoryRoutes.GET("/", controller.GetAllCategories)
}
