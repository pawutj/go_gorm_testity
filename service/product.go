package service

import (
	"github.com/pawutj/go_gorm_testity/entities"
	"github.com/pawutj/go_gorm_testity/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

// type ProductService interface {
// 	Create(product *entities.Product) error
// }

func InitialProductService(repository repository.ProductRepository) ProductService {
	productService := ProductService{productRepository: repository}
	return productService
}

func (productService ProductService) Create(product *entities.Product) error {
	err := productService.productRepository.Create(product)
	if err != nil {

	}
	return nil

}
