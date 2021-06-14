// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	apifarm "apifarm/src"

	mock "github.com/stretchr/testify/mock"
)

// QueryFactory is an autogenerated mock type for the QueryFactory type
type QueryFactory struct {
	mock.Mock
}

// Build provides a mock function with given fields: _a0, _a1
func (_m *QueryFactory) Build(_a0 []byte, _a1 uint) apifarm.Query {
	ret := _m.Called(_a0, _a1)

	var r0 apifarm.Query
	if rf, ok := ret.Get(0).(func([]byte, uint) apifarm.Query); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(apifarm.Query)
	}

	return r0
}

// Error provides a mock function with given fields: _a0
func (_m *QueryFactory) Error(_a0 error) apifarm.Query {
	ret := _m.Called(_a0)

	var r0 apifarm.Query
	if rf, ok := ret.Get(0).(func(error) apifarm.Query); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(apifarm.Query)
	}

	return r0
}
