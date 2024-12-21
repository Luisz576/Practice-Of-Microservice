package usecase

import (
	"encoding/json"
	"luisz/core/entity"
	"microservice-core/config"
	"microservice-core/internal/dto"
	"net/http"
)

type getAllCategories struct{}

var GetAllCategories = getAllCategories{}

type CategoriesAPIGetAllResponse struct {
	Success    bool              `json:"success" binding:"required"`
	Categories []dto.CategoryDTO `json:"categories" binding:"required"`
}

func (u *getAllCategories) Execute() (*[]*entity.Category, error) {
	response, err := http.Get(
		config.GetCategoryEndpoint(config.CategoryApiEndpointGetAll),
	)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data CategoriesAPIGetAllResponse
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	categories := make([]*entity.Category, 0)
	for i := 0; i < len(data.Categories); i++ {
		categories = append(categories, &entity.Category{
			Id:        data.Categories[i].Id,
			Name:      data.Categories[i].Name,
			CreatedAt: data.Categories[i].CreatedAt,
			UpdatedAt: data.Categories[i].UpdatedAt,
		})
	}
	return &categories, nil
}
