package repository

import (
	"testing"

	config "github.com/pawutj/go_gorm_testity/config"
	"github.com/pawutj/go_gorm_testity/entities"
	"github.com/pawutj/go_gorm_testity/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type cartRepositorySuit struct {
	suite.Suite
	repository        CartRepository
	productRepository ProductRepository
	cleanupExecutor   utils.TruncateTableExecutor
}

func (suite *cartRepositorySuit) SetupSuite() {
	db := config.ConnectDB()
	repository := InitialCartRepository(db)
	productRepository := InitialProductRepository(db)
	suite.repository = repository
	suite.productRepository = productRepository
	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)
}

func (suite *cartRepositorySuit) TearDownTest() {
	// defer suite.cleanupExecutor.TruncateTable("carts_products")
	defer suite.cleanupExecutor.TruncateTable("carts")
}

func (suite *cartRepositorySuit) TestCreateCart_Positive() {

	cart := entities.Cart{
		UserId: 0,
	}
	err := suite.repository.Create(&cart)

	assert.Equal(suite.T(), err, nil)

}

func (suite *cartRepositorySuit) TestGetAll() {

	cart := entities.Cart{
		UserId: 1,
	}
	suite.repository.Create(&cart)
	carts := suite.repository.GetAll()

	assert.Greater(suite.T(), len(carts), 0)
}

func (suite *cartRepositorySuit) TestGetById() {

	_cart := entities.Cart{
		UserId: 1,
	}
	suite.repository.Create(&_cart)
	carts := suite.repository.GetAll()
	cart := carts[0]
	result := suite.repository.GetById(carts[0].ID)
	assert.Equal(suite.T(), cart.ID, result.ID)
	assert.Equal(suite.T(), cart.UserId, result.UserId)
}

func (suite *cartRepositorySuit) TestUpdate() {
	cart := entities.Cart{UserId: 20}
	err := suite.repository.Update(2, &cart)

	assert.Equal(suite.T(), err, nil)
}

// func (suite *cartRepositorySuit) TestAddProductToCart() {
// 	product := entities.Product{ProductName: "testCart", Price: 213}
// 	err := suite.repository.AddProductToCart(1, &product)
// 	assert.Equal(suite.T(), err, nil)
// }

// func (suite *cartRepositorySuit) TestAddProductToCartByProductId() {
// 	product := suite.productRepository.GetById(1)

// 	err := suite.repository.AddProductToCart(2, &product)

// 	assert.Equal(suite.T(), err, nil)

// }

func TestCartRepository(t *testing.T) {
	suite.Run(t, new(cartRepositorySuit))
}
