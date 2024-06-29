package mock_db

import (
	sql "database/sql"

	gomock "github.com/golang/mock/gomock"
)

type MySQL struct {
	DB      *sql.DB
	QUERIES *MockQuerier
}

func NewMySQLMock(ctrl *gomock.Controller) (*MySQL, error) {
	mySQL := &MySQL{}
	mySQL.QUERIES = NewMockQuerier(ctrl)

	return mySQL, nil
}
