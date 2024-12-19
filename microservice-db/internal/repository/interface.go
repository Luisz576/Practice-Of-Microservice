package repository

import "go-categories-microservice/internal/entity"

type ICategoryRepository interface {
	Save(category *entity.Category) error
	GetAll() ([]*entity.Category, error)
}
