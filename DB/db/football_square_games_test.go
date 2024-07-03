package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FootBallSquareGamesTestSuite struct {
	suite.Suite
}

func (suite *FootBallSquareGamesTestSuite) SetupTest() {}

func (suite *FootBallSquareGamesTestSuite) TestCreateFootBallSquareGames() {

}

func TestFootBallSquareGamesTestSuite(t *testing.T) {
	suite.Run(t, new(FootBallSquareGamesTestSuite))
}
