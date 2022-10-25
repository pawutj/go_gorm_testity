package repository

import (
	"github.com/pawutj/go_gorm_testity/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	Create(product *entities.Product) error
	GetAll() []entities.Product
}

func InitialProductRepository() ProductRepository {
	dsn := "root:123456@tcp(127.0.0.1:3306)/finalproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&entities.Product{})
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
