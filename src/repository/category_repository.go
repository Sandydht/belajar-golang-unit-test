package repository

import "github.com/Sandydht/belajar-golang-unit-test/src/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
