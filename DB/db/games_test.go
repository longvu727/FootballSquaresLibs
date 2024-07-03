package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GamesTestSuite struct {
	suite.Suite
}

func (suite *GamesTestSuite) SetupTest() {}

func (suite *GamesTestSuite) TestCreateGames() {

}

func TestGamesTestSuite(t *testing.T) {
	suite.Run(t, new(GamesTestSuite))
}
