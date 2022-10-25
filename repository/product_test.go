package repository

import (
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
