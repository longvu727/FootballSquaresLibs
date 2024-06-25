package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/longvu727/FootballSquaresLibs/util"
)

type MySQL struct {
	DB      *sql.DB
	QUERIES *Queries
}

func NewMySQL(config util.Config) (*MySQL, error) {
	mySQL := &MySQL{}

	db, err := sql.Open("mysql", config.MySQLDSN)
	if err != nil {
		return mySQL, err
	}

	mySQL.DB = db
	mySQL.QUERIES = New(db)

	return mySQL, err
}
