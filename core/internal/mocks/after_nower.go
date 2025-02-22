// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// AfterNower is an autogenerated mock type for the AfterNower type
type AfterNower struct {
	mock.Mock
}

// After provides a mock function with given fields: d
func (_m *AfterNower) After(d time.Duration) <-chan time.Time {
	ret := _m.Called(d)

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func(time.Duration) <-chan time.Time); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// Now provides a mock function with given fields:
func (_m *AfterNower) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}
