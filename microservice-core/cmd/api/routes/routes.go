package routes

import (
	"microservice-core/internal/controller"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/", controller.CreateCategory)
	categoryRoutes.GET("/", controller.GetAllCategories)
}
