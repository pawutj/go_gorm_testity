package entities

type Product struct {
	ProductName string `db:"product_name"`
	Price       int    `db:"price"`
}