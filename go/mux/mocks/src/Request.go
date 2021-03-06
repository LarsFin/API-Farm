// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Request is an autogenerated mock type for the Request type
type Request struct {
	mock.Mock
}

// GetBody provides a mock function with given fields:
func (_m *Request) GetBody() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParam provides a mock function with given fields: _a0
func (_m *Request) GetParam(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
