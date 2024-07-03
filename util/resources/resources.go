package resources

import (
	"context"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util"
)

type Resources struct {
	Config  util.Config
	DB      db.MySQL
	Context context.Context
}

func NewResources(config util.Config, mySQL db.MySQL, context context.Context) *Resources {
	return &Resources{
		Config:  config,
		DB:      mySQL,
		Context: context,
	}
}
