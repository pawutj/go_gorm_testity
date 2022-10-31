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

func TestCartRepository(t *testing.T) {
	suite.Run(t, new(cartRepositorySuit))
}
