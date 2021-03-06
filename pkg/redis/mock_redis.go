package redis

import (
	"context"
	"time"

	mock "github.com/stretchr/testify/mock"
)

// MockRedisStore is an autogenerated mock type for the RedisStorage type
type MockRedisStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: context, key
func (_m *MockRedisStore) Get(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeys provides a mock function with given fields: context, pattern
func (_m *MockRedisStore) GetKeys(ctx context.Context, pattern string) ([]string, error) {
	ret := _m.Called(ctx, pattern)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, pattern)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pattern)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, key, value, expire
func (_m *MockRedisStore) Set(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	ret := _m.Called(ctx, key, value, expire)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, expire)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Error(1)
		}
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, key
func (_m *MockRedisStore) Exists(ctx context.Context, key string) (bool, error) {
	ret := _m.Called(ctx, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r1 = ret.Error(1)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, key
func (_m *MockRedisStore) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Error(1)
		}
	}

	return r0
}
