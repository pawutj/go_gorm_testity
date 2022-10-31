package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID       int       `gorm:"primary_key;auto_increment"`
	UserId   int       `db:"user_id"`
	Products []Product `gorm:"many2many:carts_products;"`
}
