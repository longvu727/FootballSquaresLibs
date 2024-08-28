package db

import (
	"testing"

	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/stretchr/testify/suite"
)

type MySQLTestSuite struct {
	suite.Suite
	config util.Config
}

func TestMySQLTestSuite(t *testing.T) {
	suite.Run(t, new(MySQLTestSuite))
}

func (suite *MySQLTestSuite) TestNewMySQL() {
	suite.config = util.Config{
		MySQLDSN: "127.0.0.1",
	}
	_, err := NewMySQL(suite.config)

	suite.Error(err)
}
