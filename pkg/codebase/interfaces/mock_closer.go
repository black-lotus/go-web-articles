// Code generated by mockery 2.9.0. DO NOT EDIT.

package interfaces

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockCloser is an autogenerated mock type for the Closer type
type MockCloser struct {
	mock.Mock
}

// Disconnect provides a mock function with given fields: ctx
func (_m *MockCloser) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
