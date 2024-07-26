package footballsquaregameservices

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateFootballSquareGameServiceTestSuite struct {
	suite.Suite
}

func TestCreateFootballSquareGameServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CreateFootballSquareGameServiceTestSuite))
}

func (suite *CreateFootballSquareGameServiceTestSuite) TestRequest() {
}
