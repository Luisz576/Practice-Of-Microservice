package usecase

import (
	"go-categories-microservice/internal/repository"
	"luisz/core/entity"
)

type createCategoryUseCase struct {
	repository repository.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repository.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{
		repository: repository,
	}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entity.NewCategory(name)
	if err != nil {
		return err
	}

	err = u.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}
