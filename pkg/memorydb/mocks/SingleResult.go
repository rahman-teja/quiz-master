// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SingleResult is an autogenerated mock type for the SingleResult type
type SingleResult struct {
	mock.Mock
}

// Decode provides a mock function with given fields: dt
func (_m *SingleResult) Decode(dt interface{}) error {
	ret := _m.Called(dt)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Err provides a mock function with given fields:
func (_m *SingleResult) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
