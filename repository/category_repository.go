package repository

import "mock-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
