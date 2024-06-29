package mock_db

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/longvu727/FootballSquaresLibs/DB/db"
)

func NewMockMySQL(ctrl *gomock.Controller) (*db.MySQL, error) {
	mySQL := &db.MySQL{}

	mySQL.QUERIES = db.New(NewMockDBTX(ctrl))

	return mySQL, nil
}
