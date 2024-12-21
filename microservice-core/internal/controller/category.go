package controller

import (
	"microservice-core/internal/dto"
	"microservice-core/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	var categoryDTO dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&categoryDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	status, err := usecase.CreateNewCategory.Execute(&categoryDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(status, gin.H{
		"success": status == http.StatusCreated,
	})
}

func GetAllCategories(ctx *gin.Context) {
	categories, err := usecase.GetAllCategories.Execute()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":    true,
		"categories": categories,
	})
}
