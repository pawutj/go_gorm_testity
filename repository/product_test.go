package repository

import (
	"testing"

	config "github.com/pawutj/go_gorm_testity/config"
	"github.com/pawutj/go_gorm_testity/entities"
	"github.com/pawutj/go_gorm_testity/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type productRepositorySuit struct {
	suite.Suite
	repository      ProductRepository
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *productRepositorySuit) SetupSuite() {
	db := config.ConnectDB()

	repository := InitialProductRepository(db)
	suite.repository = repository
	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)
}

func (suite *productRepositorySuit) TearDownTest() {
	defer suite.cleanupExecutor.TruncateTable()
}

func (suite *productRepositorySuit) TestCreateProduct_Positive() {

	product := entities.Product{
		ProductName: "testCreate",
		Price:       1,
	}
	suite.repository.Create(&product)

	result := suite.repository.GetByName("testCreate")
	assert.Equal(suite.T(), result[0].Price, 1)
}

func (suite *productRepositorySuit) TestGetAll() {

	product := entities.Product{
		ProductName: "testGetAll",
		Price:       2,
	}

	suite.repository.Create(&product)

	products := suite.repository.GetAll()
	assert.Greater(suite.T(), len(products), 0)
}

func (suite *productRepositorySuit) TestGetById() {

	_product := entities.Product{
		ProductName: "testGetById",
		Price:       2,
	}
	suite.repository.Create(&_product)

	products := suite.repository.GetAll()
	product := products[0]

	result := suite.repository.GetById(product.ID)
	assert.Equal(suite.T(), result.ID, product.ID)
	assert.Equal(suite.T(), result.Price, product.Price)
	assert.Equal(suite.T(), result.ProductName, product.ProductName)

}

func (suite *productRepositorySuit) TestGetByName() {

	product := entities.Product{
		ProductName: "testGetByName",
		Price:       3,
	}
	suite.repository.Create(&product)

	products := suite.repository.GetByName("testGetByName")

	assert.GreaterOrEqual(suite.T(), len(products), 0)
	assert.Equal(suite.T(), products[0].Price, 3)
}

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(productRepositorySuit))

}
