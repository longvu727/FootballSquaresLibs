package mock_db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	gomock "github.com/golang/mock/gomock"
)

type MySQLMock struct {
	DB      *sql.DB
	QUERIES *MockQuerier
}

func NewMySQLMock(ctrl *gomock.Controller) (*MySQLMock, error) {
	mySQL := &MySQLMock{}
	mySQL.QUERIES = NewMockQuerier(ctrl)

	return mySQL, nil
}
