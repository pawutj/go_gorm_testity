package repository

import (
	"testing"

	"github.com/pawutj/go_gorm_testity/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type cartRepositorySuit struct {
	suite.Suite
	repository        CartRepository
	productRepository ProductRepository
}

func (suite *cartRepositorySuit) SetupSuite() {
	repository := InitialCartRepository()
	productRepository := InitialProductRepository()
	suite.repository = repository
	suite.productRepository = productRepository
}

func (suite *cartRepositorySuit) TestCreateCart_Positive() {

	cart := entities.Cart{
		UserId: 0,
	}
	err := suite.repository.Create(&cart)

	assert.Equal(suite.T(), err, nil)

}

func (suite *cartRepositorySuit) TestGetById() {
	cart := suite.repository.GetById(0)
	assert.Equal(suite.T(), cart.ID, 0)
	assert.Equal(suite.T(), cart.UserId, 0)
}

func (suite *cartRepositorySuit) TestAddProductToCart() {
	product := entities.Product{ProductName: "testCart", Price: 213}
	err := suite.repository.AddProductToCart(1, &product)
	assert.Equal(suite.T(), err, nil)
}

func (suite *cartRepositorySuit) TestUpdate() {
	cart := entities.Cart{UserId: 20}
	err := suite.repository.Update(2, &cart)

	assert.Equal(suite.T(), err, nil)
}

func (suite *cartRepositorySuit) TestAddProductToCartByProductId() {
	product := suite.productRepository.GetById(1)

	err := suite.repository.AddProductToCart(2, &product)

	assert.Equal(suite.T(), err, nil)

}

func TestCartRepository(t *testing.T) {
	suite.Run(t, new(cartRepositorySuit))
}
