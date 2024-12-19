package routes

import (
	"go-categories-microservice/cmd/api/controller"
	"go-categories-microservice/internal/repository"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRepository := repository.NewInMemoryCategoryRepository()

	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controller.CreateCategory(ctx, categoryRepository)
	})
	categoryRoutes.GET("/", func(ctx *gin.Context) {
		controller.GetAllCategories(ctx, categoryRepository)
	})
}
