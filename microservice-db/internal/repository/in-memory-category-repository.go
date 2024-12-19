package repository

import "go-categories-microservice/internal/entity"

type inMemoryCategoryRepository struct {
	db []*entity.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entity.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entity.Category) error {
	r.db = append(r.db, category)
	return nil
}

func (r *inMemoryCategoryRepository) GetAll() ([]*entity.Category, error) {
	return r.db, nil
}
