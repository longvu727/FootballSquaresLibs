package util

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (suite *ConfigTestSuite) TestLoadConfig() {
	config, err := LoadConfig(".", "app_test", "env")
	if err != nil {
		suite.Fail("Unable to load config file, " + err.Error())
	}

	suite.Equal("123", config.MySQLDSN)
	suite.Equal("1000", config.PORT)
	suite.Equal("redis_url", config.REDISURL)
}

func (suite *ConfigTestSuite) TestLoadConfigJSON() {
	config, err := LoadConfig(".", "app_json_test", "json")
	if err != nil {
		suite.Fail("Unable to load config file, " + err.Error())
	}

	suite.Equal("1000", config.PORT)
	suite.Len(config.MICROSERVICESBASEURL, 4)
	suite.Equal("http://squaremicroservices:3000", config.MICROSERVICESBASEURL["squaremicroservices"])
}

func TestConfigTestSuite(T *testing.T) {
	suite.Run(T, new(ConfigTestSuite))
}
