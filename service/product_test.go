package service

import (
	"github.com/pawutj/go_gorm_testity/entities"
	"github.com/pawutj/go_gorm_testity/mock/mock_repository"
	"github.com/stretchr/testify/suite"
)

type ProductServiceSuite struct {
	suite.Suite
	repository *mock_repository.ProductRepository
	service    ProductService
}

func (suite *ProductServiceSuite) SetupTest() {
	repository := new(mock_repository.ProductRepository)
	service := InitialProductService(repository)

	suite.repository = repository
	suite.service = service
}

func (suite *ProductServiceSuite) TestCreate_Positive() {
	product := entities.Product{}
	suite.repository.On("Create", &product).Return(nil)

	err := suite.service.Create(&product)

	suite.Nil(err, "pass")

}
