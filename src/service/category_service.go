package service

import (
	"errors"

	"github.com/Sandydht/belajar-golang-unit-test/src/entity"
	"github.com/Sandydht/belajar-golang-unit-test/src/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}
