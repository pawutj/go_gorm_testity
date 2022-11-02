package repository

import (
	"testing"

	"github.com/pawutj/go_gorm_testity/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type cartRepositorySuit struct {
	suite.Suite
	repository CartRepository
}

func (suite *cartRepositorySuit) SetupSuite() {
	repository := InitialCartRepository()
	suite.repository = repository
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
	product := entities.Product{ProductName: "testCart", Price: 200}
	err := suite.repository.AddProductToCart(1, &product)
	assert.Equal(suite.T(), err, nil)
}

func TestCartRepository(t *testing.T) {
	suite.Run(t, new(cartRepositorySuit))
}
