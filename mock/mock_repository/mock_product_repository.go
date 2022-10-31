package mock_repository

import (
	"github.com/pawutj/go_gorm_testity/entities"
	mock "github.com/stretchr/testify/mock"
)

type ProductRepository struct {
	mock.Mock
}

func (m *ProductRepository) Create(product *entities.Product) error {
	args := m.Called(product)

	return args.Error(0)
}

func (m *ProductRepository) GetAll() []entities.Product {
	products := []entities.Product{}

	return products
}

func (m *ProductRepository) GetById(id int) entities.Product {
	product := entities.Product{}

	return product
}

func (m *ProductRepository) GetByName(name string) []entities.Product {
	product := []entities.Product{}
	return product
}
