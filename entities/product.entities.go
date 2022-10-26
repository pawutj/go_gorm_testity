package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          int    `gorm:"primary_key;auto_increment"`
	ProductName string `db:"product_name"`
	Price       int    `db:"price"`
}
