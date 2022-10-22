package mock_repository

import (
	"github.com/pawutj/go_gorm_testity/entities"
	mock "github.com/stretchr/testify/mock"
)

type ProductRepository struct {
	mock.Mock
}

func (m *ProductRepository) Create(product *entities.Product) error {
	ret := m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
