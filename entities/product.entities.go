package entities

type Product struct {
	ID          int    `gorm:"primary_key;auto_increment"`
	ProductName string `db:"product_name"`
	Price       int    `db:"price"`
}
