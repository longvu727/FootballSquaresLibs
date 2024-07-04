// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/longvu727/FootballSquaresLibs/DB/db (interfaces: MySQL)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/longvu727/FootballSquaresLibs/DB/db"
)

// MockMySQL is a mock of MySQL interface.
type MockMySQL struct {
	ctrl     *gomock.Controller
	recorder *MockMySQLMockRecorder
}

// MockMySQLMockRecorder is the mock recorder for MockMySQL.
type MockMySQLMockRecorder struct {
	mock *MockMySQL
}

// NewMockMySQL creates a new mock instance.
func NewMockMySQL(ctrl *gomock.Controller) *MockMySQL {
	mock := &MockMySQL{ctrl: ctrl}
	mock.recorder = &MockMySQLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMySQL) EXPECT() *MockMySQLMockRecorder {
	return m.recorder
}

// CreateFootballSquareGame mocks base method.
func (m *MockMySQL) CreateFootballSquareGame(arg0 context.Context, arg1 db.CreateFootballSquareGameParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFootballSquareGame", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFootballSquareGame indicates an expected call of CreateFootballSquareGame.
func (mr *MockMySQLMockRecorder) CreateFootballSquareGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFootballSquareGame", reflect.TypeOf((*MockMySQL)(nil).CreateFootballSquareGame), arg0, arg1)
}

// CreateGame mocks base method.
func (m *MockMySQL) CreateGame(arg0 context.Context, arg1 db.CreateGameParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGame", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGame indicates an expected call of CreateGame.
func (mr *MockMySQLMockRecorder) CreateGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGame", reflect.TypeOf((*MockMySQL)(nil).CreateGame), arg0, arg1)
}

// CreateSquare mocks base method.
func (m *MockMySQL) CreateSquare(arg0 context.Context, arg1 db.CreateSquareParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSquare", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSquare indicates an expected call of CreateSquare.
func (mr *MockMySQLMockRecorder) CreateSquare(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSquare", reflect.TypeOf((*MockMySQL)(nil).CreateSquare), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockMySQL) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockMySQLMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockMySQL)(nil).CreateUser), arg0, arg1)
}

// GetFootballSquareGame mocks base method.
func (m *MockMySQL) GetFootballSquareGame(arg0 context.Context, arg1 int32) (db.GetFootballSquareGameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFootballSquareGame", arg0, arg1)
	ret0, _ := ret[0].(db.GetFootballSquareGameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFootballSquareGame indicates an expected call of GetFootballSquareGame.
func (mr *MockMySQLMockRecorder) GetFootballSquareGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFootballSquareGame", reflect.TypeOf((*MockMySQL)(nil).GetFootballSquareGame), arg0, arg1)
}

// GetFootballSquareGameByGameID mocks base method.
func (m *MockMySQL) GetFootballSquareGameByGameID(arg0 context.Context, arg1 sql.NullInt32) ([]db.GetFootballSquareGameByGameIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFootballSquareGameByGameID", arg0, arg1)
	ret0, _ := ret[0].([]db.GetFootballSquareGameByGameIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFootballSquareGameByGameID indicates an expected call of GetFootballSquareGameByGameID.
func (mr *MockMySQLMockRecorder) GetFootballSquareGameByGameID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFootballSquareGameByGameID", reflect.TypeOf((*MockMySQL)(nil).GetFootballSquareGameByGameID), arg0, arg1)
}

// GetGame mocks base method.
func (m *MockMySQL) GetGame(arg0 context.Context, arg1 int32) (db.GetGameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGame", arg0, arg1)
	ret0, _ := ret[0].(db.GetGameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGame indicates an expected call of GetGame.
func (mr *MockMySQLMockRecorder) GetGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGame", reflect.TypeOf((*MockMySQL)(nil).GetGame), arg0, arg1)
}

// GetGameByGUID mocks base method.
func (m *MockMySQL) GetGameByGUID(arg0 context.Context, arg1 string) (db.GetGameByGUIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameByGUID", arg0, arg1)
	ret0, _ := ret[0].(db.GetGameByGUIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameByGUID indicates an expected call of GetGameByGUID.
func (mr *MockMySQLMockRecorder) GetGameByGUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameByGUID", reflect.TypeOf((*MockMySQL)(nil).GetGameByGUID), arg0, arg1)
}

// GetSquare mocks base method.
func (m *MockMySQL) GetSquare(arg0 context.Context, arg1 int32) (db.GetSquareRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSquare", arg0, arg1)
	ret0, _ := ret[0].(db.GetSquareRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSquare indicates an expected call of GetSquare.
func (mr *MockMySQLMockRecorder) GetSquare(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSquare", reflect.TypeOf((*MockMySQL)(nil).GetSquare), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockMySQL) GetUser(arg0 context.Context, arg1 int32) (db.GetUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.GetUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockMySQLMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockMySQL)(nil).GetUser), arg0, arg1)
}
