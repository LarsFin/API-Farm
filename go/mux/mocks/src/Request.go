// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Request is an autogenerated mock type for the Request type
type Request struct {
	mock.Mock
}

// GetBody provides a mock function with given fields:
func (_m *Request) GetBody() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}
