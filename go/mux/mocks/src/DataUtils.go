// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	apifarm "apifarm/src"

	mock "github.com/stretchr/testify/mock"
)

// DataUtils is an autogenerated mock type for the DataUtils type
type DataUtils struct {
	mock.Mock
}

// DeserializeConfiguration provides a mock function with given fields: _a0
func (_m *DataUtils) DeserializeConfiguration(_a0 []byte) (*apifarm.Configuration, error) {
	ret := _m.Called(_a0)

	var r0 *apifarm.Configuration
	if rf, ok := ret.Get(0).(func([]byte) *apifarm.Configuration); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apifarm.Configuration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeserializeVideoGame provides a mock function with given fields: _a0
func (_m *DataUtils) DeserializeVideoGame(_a0 []byte) (*apifarm.VideoGame, error) {
	ret := _m.Called(_a0)

	var r0 *apifarm.VideoGame
	if rf, ok := ret.Get(0).(func([]byte) *apifarm.VideoGame); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apifarm.VideoGame)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeserializeVideoGames provides a mock function with given fields: _a0
func (_m *DataUtils) DeserializeVideoGames(_a0 []byte) (*[]apifarm.VideoGame, error) {
	ret := _m.Called(_a0)

	var r0 *[]apifarm.VideoGame
	if rf, ok := ret.Get(0).(func([]byte) *[]apifarm.VideoGame); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]apifarm.VideoGame)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Serialize provides a mock function with given fields: _a0
func (_m *DataUtils) Serialize(_a0 interface{}) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
