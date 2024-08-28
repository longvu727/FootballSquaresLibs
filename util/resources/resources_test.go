package resources

import (
	"context"
	"testing"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/stretchr/testify/suite"
)

type ResourcesTestSuite struct {
	suite.Suite
}

func TestResourcesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourcesTestSuite))
}

func (suite *ResourcesTestSuite) TestNewResources() {
	config := util.Config{}
	mySQL := db.MySQLQueries{}
	services := services.NewServices()
	context := context.Background()

	R := NewResources(config, mySQL, services, context)

	suite.Equal(config, R.Config)
	suite.Equal(mySQL, R.DB)
	suite.Equal(services, R.Services)
	suite.Equal(context, R.Context)
}
