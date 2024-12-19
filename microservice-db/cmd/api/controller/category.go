package controller

import (
	"go-categories-microservice/internal/repository"
	"go-categories-microservice/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(context *gin.Context, repository repository.ICategoryRepository) {
	var body createCategoryInput

	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	useCase := usecase.NewCreateCategoryUseCase(repository)
	err := useCase.Execute(body.Name)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": true})
}

func GetAllCategories(context *gin.Context, repository repository.ICategoryRepository) {
	useCase := usecase.NewListCategoryUseCase(repository)
	categories, err := useCase.Execute()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success":    true,
		"categories": categories,
	})
}
