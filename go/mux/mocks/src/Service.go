// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	apifarm "apifarm/src"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *Service) GetAll() apifarm.Query {
	ret := _m.Called()

	var r0 apifarm.Query
	if rf, ok := ret.Get(0).(func() apifarm.Query); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(apifarm.Query)
	}

	return r0
}
