package repository

import (
	"testing"

	"github.com/pawutj/go_gorm_testity/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type productRepositorySuit struct {
	suite.Suite
	repository ProductRepository
}

func (suite *productRepositorySuit) SetupSuite() {
	repository := InitialProductRepository()
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

// func (suite *productRepositorySuit) TestFindById

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(productRepositorySuit))
}
