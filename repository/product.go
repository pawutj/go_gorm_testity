package repository

type Product struct {
	ProductName string `db:"product_name"`
	Price       int    `db:"price"`
}

type ProductRepository interface {
	Create(Product) (*Product, error)
	GetAll() ([]Product, error)
}