package repository

import (
	"mock-testing/entity"

	"github.com/stretchr/testify/mock"
)

// membuat moock
type CategoryRepositoryMock struct {
	Mock mock.Mock
}

// implementasi kontraks
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	//memanggil program mock
	arguments := repository.Mock.Called(id) //called harus mengembalikan parameter

	//pengecekan untuk return pointer dan error
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category) //return 0 adalah interface
		return &category
	}
}
