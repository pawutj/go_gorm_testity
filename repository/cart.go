package repository

import (
	"github.com/pawutj/go_gorm_testity/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

type CartRepository interface {
	Create(cart *entities.Cart) error
	GetById(id int) entities.Cart
	AddProductToCart(cartId int, product *entities.Product) error
	Update(cartId int, cart *entities.Cart) error
}

func InitialCartRepository() CartRepository {
	dsn := "admin:123456@tcp(127.0.0.1:3307)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&entities.Cart{})
	repository := cartRepository{}
	repository.db = db
	return &repository
}

func (repository *cartRepository) Create(cart *entities.Cart) error {

	dbc := repository.db.Create(&cart)
	if dbc != nil {
		return dbc.Error
	}

	return nil
}

func (repository *cartRepository) GetById(id int) entities.Cart {
	cart := entities.Cart{}
	repository.db.First(&cart, "id = ?", id)
	return cart
}

func (repository *cartRepository) AddProductToCart(cartId int, product *entities.Product) error {
	cart := repository.GetById(cartId)
	// products := []entities.Product{{ProductName: product.ProductName, Price: product.Price}}
	cart.Products = append(cart.Products, *product)
	dbc := repository.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&cart)
	if dbc != nil {
		return dbc.Error
	}
	return nil
}

func (repository *cartRepository) Update(cartId int, cart *entities.Cart) error {
	_cart := repository.GetById(cartId)
	_cart.UserId = cart.UserId
	dbc := repository.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&_cart)
	if dbc != nil {
		return dbc.Error
	}
	return nil
}
