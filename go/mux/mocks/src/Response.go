// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Response is an autogenerated mock type for the Response type
type Response struct {
	mock.Mock
}

// BadRequestText provides a mock function with given fields: _a0
func (_m *Response) BadRequestText(_a0 string) {
	_m.Called(_a0)
}

// CreatedJSON provides a mock function with given fields: _a0
func (_m *Response) CreatedJSON(_a0 []byte) {
	_m.Called(_a0)
}

// Error provides a mock function with given fields: _a0
func (_m *Response) Error(_a0 error) {
	_m.Called(_a0)
}

// NotFoundText provides a mock function with given fields: _a0
func (_m *Response) NotFoundText(_a0 string) {
	_m.Called(_a0)
}

// OkJSON provides a mock function with given fields: _a0
func (_m *Response) OkJSON(_a0 []byte) {
	_m.Called(_a0)
}

// OkText provides a mock function with given fields: _a0
func (_m *Response) OkText(_a0 string) {
	_m.Called(_a0)
}
