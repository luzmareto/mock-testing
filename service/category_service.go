package service

import (
	"errors"
	"mock-testing/entity"
	"mock-testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id) //id adalah parameter dari get dan FindById

	//melakukan pengecekan untuk return pointer dan error
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
