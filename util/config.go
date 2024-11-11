package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	REDISURL             string            `mapstructure:"REDIS_URL"`
	MySQLDSN             string            `mapstructure:"MYSQL_DSN"`
	PORT                 string            `mapstructure:"PORT"`
	MICROSERVICESBASEURL map[string]string `mapstructure:"microservices"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, configName string, configType string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

func LoadCustomConfig(path string, configName string, configType string, config interface{}) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(config)
	return err
}
