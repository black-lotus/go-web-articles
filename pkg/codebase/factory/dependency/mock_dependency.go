// Code generated by mockery 2.9.0. DO NOT EDIT.

package dependency

import (
	interfaces "webarticles/pkg/codebase/interfaces"

	mock "github.com/stretchr/testify/mock"
)

// MockDependency is an autogenerated mock type for the Dependency type
type MockDependency struct {
	mock.Mock
}

// GetRedisPool provides a mock function with given fields:
func (_m *MockDependency) GetRedisPool() interfaces.RedisPool {
	ret := _m.Called()

	var r0 interfaces.RedisPool
	if rf, ok := ret.Get(0).(func() interfaces.RedisPool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.RedisPool)
		}
	}

	return r0
}

// GetSQLDatabase provides a mock function with given fields:
func (_m *MockDependency) GetSQLDatabase() interfaces.SQLDatabase {
	ret := _m.Called()

	var r0 interfaces.SQLDatabase
	if rf, ok := ret.Get(0).(func() interfaces.SQLDatabase); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.SQLDatabase)
		}
	}

	return r0
}

// GetValidator provides a mock function with given fields:
func (_m *MockDependency) GetValidator() interfaces.Validator {
	ret := _m.Called()

	var r0 interfaces.Validator
	if rf, ok := ret.Get(0).(func() interfaces.Validator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.Validator)
		}
	}

	return r0
}
