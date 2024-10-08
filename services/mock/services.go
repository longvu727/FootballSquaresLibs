// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/longvu727/FootballSquaresLibs/services (interfaces: Services)

// Package mockservices is a generated GoMock package.
package mockservices

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	services "github.com/longvu727/FootballSquaresLibs/services"
	util "github.com/longvu727/FootballSquaresLibs/util"
)

// MockServices is a mock of Services interface.
type MockServices struct {
	ctrl     *gomock.Controller
	recorder *MockServicesMockRecorder
}

// MockServicesMockRecorder is the mock recorder for MockServices.
type MockServicesMockRecorder struct {
	mock *MockServices
}

// NewMockServices creates a new mock instance.
func NewMockServices(ctrl *gomock.Controller) *MockServices {
	mock := &MockServices{ctrl: ctrl}
	mock.recorder = &MockServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServices) EXPECT() *MockServicesMockRecorder {
	return m.recorder
}

// CreateFootballSquareGame mocks base method.
func (m *MockServices) CreateFootballSquareGame(arg0 *util.Config, arg1 services.CreateFootballSquareGameRequest) (services.CreateFootballSquareGameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFootballSquareGame", arg0, arg1)
	ret0, _ := ret[0].(services.CreateFootballSquareGameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFootballSquareGame indicates an expected call of CreateFootballSquareGame.
func (mr *MockServicesMockRecorder) CreateFootballSquareGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFootballSquareGame", reflect.TypeOf((*MockServices)(nil).CreateFootballSquareGame), arg0, arg1)
}

// CreateGame mocks base method.
func (m *MockServices) CreateGame(arg0 *util.Config, arg1 services.CreateGameRequest) (services.CreateGameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGame", arg0, arg1)
	ret0, _ := ret[0].(services.CreateGameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGame indicates an expected call of CreateGame.
func (mr *MockServicesMockRecorder) CreateGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGame", reflect.TypeOf((*MockServices)(nil).CreateGame), arg0, arg1)
}

// CreateSquare mocks base method.
func (m *MockServices) CreateSquare(arg0 *util.Config, arg1 services.CreateSquareRequest) (services.CreateSquareResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSquare", arg0, arg1)
	ret0, _ := ret[0].(services.CreateSquareResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSquare indicates an expected call of CreateSquare.
func (mr *MockServicesMockRecorder) CreateSquare(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSquare", reflect.TypeOf((*MockServices)(nil).CreateSquare), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockServices) CreateUser(arg0 *util.Config, arg1 services.CreateUserRequest) (services.CreateUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(services.CreateUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServicesMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockServices)(nil).CreateUser), arg0, arg1)
}

// GetFootballSquareGameByGameID mocks base method.
func (m *MockServices) GetFootballSquareGameByGameID(arg0 *util.Config, arg1 services.GetFootballSquareGameByGameIDRequest) (services.GetFootballSquareGameByGameIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFootballSquareGameByGameID", arg0, arg1)
	ret0, _ := ret[0].(services.GetFootballSquareGameByGameIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFootballSquareGameByGameID indicates an expected call of GetFootballSquareGameByGameID.
func (mr *MockServicesMockRecorder) GetFootballSquareGameByGameID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFootballSquareGameByGameID", reflect.TypeOf((*MockServices)(nil).GetFootballSquareGameByGameID), arg0, arg1)
}

// GetGame mocks base method.
func (m *MockServices) GetGame(arg0 *util.Config, arg1 services.GetGameRequest) (services.GetGameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGame", arg0, arg1)
	ret0, _ := ret[0].(services.GetGameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGame indicates an expected call of GetGame.
func (mr *MockServicesMockRecorder) GetGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGame", reflect.TypeOf((*MockServices)(nil).GetGame), arg0, arg1)
}

// GetGameByGUID mocks base method.
func (m *MockServices) GetGameByGUID(arg0 *util.Config, arg1 services.GetGameByGUIDRequest) (services.GetGameByGUIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameByGUID", arg0, arg1)
	ret0, _ := ret[0].(services.GetGameByGUIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameByGUID indicates an expected call of GetGameByGUID.
func (mr *MockServicesMockRecorder) GetGameByGUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameByGUID", reflect.TypeOf((*MockServices)(nil).GetGameByGUID), arg0, arg1)
}

// GetSquare mocks base method.
func (m *MockServices) GetSquare(arg0 *util.Config, arg1 services.GetSquareRequest) (services.GetSquareResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSquare", arg0, arg1)
	ret0, _ := ret[0].(services.GetSquareResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSquare indicates an expected call of GetSquare.
func (mr *MockServicesMockRecorder) GetSquare(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSquare", reflect.TypeOf((*MockServices)(nil).GetSquare), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockServices) GetUser(arg0 *util.Config, arg1 services.GetUserRequest) (services.GetUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(services.GetUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServicesMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockServices)(nil).GetUser), arg0, arg1)
}

// GetUserByGUID mocks base method.
func (m *MockServices) GetUserByGUID(arg0 *util.Config, arg1 services.GetUserByGUIDRequest) (services.GetUserByGUIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByGUID", arg0, arg1)
	ret0, _ := ret[0].(services.GetUserByGUIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByGUID indicates an expected call of GetUserByGUID.
func (mr *MockServicesMockRecorder) GetUserByGUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByGUID", reflect.TypeOf((*MockServices)(nil).GetUserByGUID), arg0, arg1)
}

// Post mocks base method.
func (m *MockServices) Post(arg0 string, arg1, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockServicesMockRecorder) Post(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockServices)(nil).Post), arg0, arg1, arg2)
}

// ReserveFootballSquare mocks base method.
func (m *MockServices) ReserveFootballSquare(arg0 *util.Config, arg1 services.ReserveFootballSquareRequest) (services.ReserveFootballSquareResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveFootballSquare", arg0, arg1)
	ret0, _ := ret[0].(services.ReserveFootballSquareResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReserveFootballSquare indicates an expected call of ReserveFootballSquare.
func (mr *MockServicesMockRecorder) ReserveFootballSquare(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveFootballSquare", reflect.TypeOf((*MockServices)(nil).ReserveFootballSquare), arg0, arg1)
}
