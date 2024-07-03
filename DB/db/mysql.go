package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/longvu727/FootballSquaresLibs/util"
)

type MySQL interface {
	Querier
}

type MySQLQueries struct {
	DB *sql.DB
	*Queries
}

func NewMySQL(config util.Config) (MySQL, error) {

	dbconn, _ := sql.Open("mysql", config.MySQLDSN)
	return &MySQLQueries{
			DB:      dbconn,
			Queries: New(dbconn),
		},
		nil
}
