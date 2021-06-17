// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	apifarm "apifarm/src"

	mock "github.com/stretchr/testify/mock"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// AddVideoGame provides a mock function with given fields: _a0
func (_m *DB) AddVideoGame(_a0 apifarm.VideoGame) apifarm.VideoGame {
	ret := _m.Called(_a0)

	var r0 apifarm.VideoGame
	if rf, ok := ret.Get(0).(func(apifarm.VideoGame) apifarm.VideoGame); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(apifarm.VideoGame)
	}

	return r0
}

// GetAllVideoGames provides a mock function with given fields:
func (_m *DB) GetAllVideoGames() []apifarm.VideoGame {
	ret := _m.Called()

	var r0 []apifarm.VideoGame
	if rf, ok := ret.Get(0).(func() []apifarm.VideoGame); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]apifarm.VideoGame)
		}
	}

	return r0
}