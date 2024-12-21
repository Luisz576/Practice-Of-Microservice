package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"microservice-core/config"
	"microservice-core/internal/dto"
	"net/http"
)

type createNewCategory struct{}

var CreateNewCategory = &createNewCategory{}

type CreateNewCategoryApiResponse struct {
	Success bool `json:"success" binding:"required"`
}

type CreateNewCategoryApiData struct {
	Name string `json:"name"`
}

func createDataToSend(category *dto.CategoryDTO) *CreateNewCategoryApiData {
	return &CreateNewCategoryApiData{
		Name: category.Name,
	}
}

func (u *createNewCategory) Execute(category *dto.CategoryDTO) (int, error) {
	dataToSend := createDataToSend(category)
	dataToSendBytes, err := json.Marshal(dataToSend)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	response, err := http.Post(
		config.GetCategoryEndpoint(config.CategoryApiEndpointCreate),
		"application/json",
		bytes.NewBuffer(dataToSendBytes),
	)
	if err != nil {
		return response.StatusCode, err
	}

	var data CreateNewCategoryApiResponse
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return response.StatusCode, nil
	}
	defer response.Body.Close()

	if !data.Success {
		return response.StatusCode, errors.New("couldn't create category")
	}

	return response.StatusCode, nil
}
