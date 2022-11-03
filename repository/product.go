package repository

import (
	"github.com/pawutj/go_gorm_testity/entities"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	Create(product *entities.Product) error
	GetAll() []entities.Product
	GetById(id int) entities.Product
	GetByName(name string) []entities.Product
}

func InitialProductRepository(db *gorm.DB) ProductRepository {

	repository := productRepository{}
	repository.db = db
	return &repository
}

func (repository *productRepository) Create(product *entities.Product) error {

	dbc := repository.db.Create(&product)
	if dbc != nil {
		return nil
	}

	return nil
}

func (repository *productRepository) GetAll() []entities.Product {
	products := []entities.Product{}
	repository.db.Find(&products)
	return products
}

func (repository *productRepository) GetById(id int) entities.Product {
	product := entities.Product{}
	repository.db.First(&product, "id = ?", id)
	return product
}

func (repository *productRepository) GetByName(name string) []entities.Product {
	products := []entities.Product{}
	repository.db.Find(&products, "product_name Like ?", name)
	return products
}
