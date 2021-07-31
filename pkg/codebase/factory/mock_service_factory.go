// Code generated by mockery 2.9.0. DO NOT EDIT.

package factory

import (
	dependency "webarticles/pkg/codebase/factory/dependency"

	mock "github.com/stretchr/testify/mock"

	types "webarticles/pkg/codebase/factory/types"
)

// MockServiceFactory is an autogenerated mock type for the ServiceFactory type
type MockServiceFactory struct {
	mock.Mock
}

// GetDependency provides a mock function with given fields:
func (_m *MockServiceFactory) GetDependency() dependency.Dependency {
	ret := _m.Called()

	var r0 dependency.Dependency
	if rf, ok := ret.Get(0).(func() dependency.Dependency); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dependency.Dependency)
		}
	}

	return r0
}

// GetModules provides a mock function with given fields:
func (_m *MockServiceFactory) GetModules() []ModuleFactory {
	ret := _m.Called()

	var r0 []ModuleFactory
	if rf, ok := ret.Get(0).(func() []ModuleFactory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ModuleFactory)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *MockServiceFactory) Name() types.Service {
	ret := _m.Called()

	var r0 types.Service
	if rf, ok := ret.Get(0).(func() types.Service); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Service)
	}

	return r0
}
