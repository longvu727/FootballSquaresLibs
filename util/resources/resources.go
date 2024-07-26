package resources

import (
	"context"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/redis/go-redis/v9"
)

type Resources struct {
	Config      util.Config
	DB          db.MySQL
	Context     context.Context
	RedisClient *redis.Client
	Services    services.Services
}

func NewResources(config util.Config, mySQL db.MySQL, services services.Services, context context.Context) *Resources {
	return &Resources{
		Config:   config,
		DB:       mySQL,
		Context:  context,
		Services: services,
	}
}
