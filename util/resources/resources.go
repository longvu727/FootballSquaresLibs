package resources

import (
	"context"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util"
)

type Resources struct {
	util.Config
	db.MySQL
	context.Context
}

func NewResources(config util.Config, mySQL db.MySQL, context context.Context) *Resources {
	return &Resources{
		Config:  config,
		MySQL:   mySQL,
		Context: context,
	}
}
