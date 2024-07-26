package footballsquaregameservices

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReserveFootballSquareServiceTestSuite struct {
	suite.Suite
}

func TestReserveFootballSquareServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ReserveFootballSquareServiceTestSuite))
}

func (suite *ReserveFootballSquareServiceTestSuite) TestRequest() {
}
