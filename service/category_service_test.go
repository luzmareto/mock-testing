package service

import (
	"mock-testing/entity"
	"mock-testing/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Memanggil program mock
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil) //simulasi connect db yang returnya nil

	//simulasi query
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccsess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	//kontrak FindById dengan value string 2 pada parameter id
	categoryRepository.Mock.On("FindById", "2").Return(category) //return wajib interface dari category

	// melakukan perbandingan
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}
