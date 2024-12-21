package repository

import "luisz/core/entity"

type ICategoryRepository interface {
	Save(category *entity.Category) error
	GetAll() ([]*entity.Category, error)
}
