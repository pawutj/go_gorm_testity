package repository

import "github.com/pawutj/go_gorm_testity/entities"

type ProductRepository interface {
	Create(product *entities.Product) error
}
