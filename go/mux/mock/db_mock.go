// Code generated by MockGen. DO NOT EDIT.
// Source: src\db.go

// Package mock_apifarm is a generated GoMock package.
package mock_apifarm

import (
	apifarm "apifarm/src"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// GetAllVideoGames mocks base method.
func (m *MockDB) GetAllVideoGames() []apifarm.VideoGame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVideoGames")
	ret0, _ := ret[0].([]apifarm.VideoGame)
	return ret0
}

// GetAllVideoGames indicates an expected call of GetAllVideoGames.
func (mr *MockDBMockRecorder) GetAllVideoGames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVideoGames", reflect.TypeOf((*MockDB)(nil).GetAllVideoGames))
}
