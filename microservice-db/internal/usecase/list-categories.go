package usecase

import (
	"go-categories-microservice/internal/entity"
	"go-categories-microservice/internal/repository"
)

type listCategoryUseCase struct {
	repository repository.ICategoryRepository
}

func NewListCategoryUseCase(repository repository.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{
		repository: repository,
	}
}

func (u *listCategoryUseCase) Execute() ([]*entity.Category, error) {
	categories, err := u.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
