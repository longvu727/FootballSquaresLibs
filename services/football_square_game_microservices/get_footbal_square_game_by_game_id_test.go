package footballsquaregameservices

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetFootballSquareGameByGameIDServiceTestSuite struct {
	suite.Suite
}

func TestGetFootballSquareGameByGameIDServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GetFootballSquareGameByGameIDServiceTestSuite))
}

func (suite *GetFootballSquareGameByGameIDServiceTestSuite) TestRequest() {
}
