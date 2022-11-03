package repository

import (
	"testing"

	config "github.com/pawutj/go_gorm_testity/config"
	"github.com/pawutj/go_gorm_testity/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type productRepositorySuit struct {
	suite.Suite
	repository ProductRepository
}

func (suite *productRepositorySuit) SetupSuite() {
	db := config.ConnectDB()

	repository := InitialProductRepository(db)
	suite.repository = repository
}

func (suite *productRepositorySuit) TestCreateProduct_Positive() {

	product := entities.Product{
		ProductName: "test",
		Price:       1,
	}
	err := suite.repository.Create(&product)
	if err != nil {

	}
	suite.NoError(err, "no error")

}

func (suite *productRepositorySuit) TestGetAll() {
	products := suite.repository.GetAll()
	assert.Greater(suite.T(), len(products), 1)
}

func (suite *productRepositorySuit) TestGetById() {
	product := suite.repository.GetById(1)
	assert.Equal(suite.T(), product.ID, 1)
	assert.Equal(suite.T(), product.ProductName, "test")
}

func (suite *productRepositorySuit) TestGetByName() {
	product := suite.repository.GetByName("test")

	assert.GreaterOrEqual(suite.T(), len(product), 1)
}

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(productRepositorySuit))
}
