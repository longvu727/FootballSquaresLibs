package resources

import (
	"context"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util"
)

type Resources struct {
	config  util.Config
	db      db.MySQL
	context context.Context
}

func NewResources(config util.Config, db db.MySQL, context context.Context) *Resources {
	return &Resources{
		config:  config,
		db:      db,
		context: context,
	}
}
