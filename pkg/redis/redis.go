package store

import (
	"context"
	"fmt"
	"time"

	"webarticles/pkg/codebase/interfaces"
	"webarticles/pkg/logger"

	"github.com/gomodule/redigo/redis"
)

// RedisStore redis
type RedisStore struct {
	read, write *redis.Pool
}

// NewRedisStore constructor
func NewRedisStore(read, write *redis.Pool) interfaces.Store {
	return &RedisStore{read: read, write: write}
}

// Get method
func (r *RedisStore) Get(ctx context.Context, key string) (string, error) {
	// set client
	cl := r.read.Get()
	defer cl.Close()

	var data string
	var err error
	data, err = redis.String(cl.Do("GET", key))
	if err != nil {
		logger.LogRed(err.Error())
		return data, err
	}

	return data, nil
}

// GetKeys method
func (r *RedisStore) GetKeys(ctx context.Context, pattern string) ([]string, error) {
	// set client
	cl := r.read.Get()
	defer cl.Close()

	var datas []string
	var err error
	datas, err = redis.Strings(cl.Do("KEYS", fmt.Sprintf("%s*", pattern)))
	if err != nil {
		logger.LogRed(err.Error())
		return datas, err
	}

	return datas, nil
}

// Set method
func (r *RedisStore) Set(ctx context.Context, key string, value interface{}, expire time.Duration) (err error) {
	// set client
	cl := r.write.Get()
	defer cl.Close()

	_, err = cl.Do("SET", key, value)
	if err != nil {
		logger.LogRed(err.Error())
		return
	}

	_, err = cl.Do("EXPIRE", key, int(expire.Seconds()))
	if err != nil {
		logger.LogRed(err.Error())
		return
	}

	return nil
}

// Exists method
func (r RedisStore) Exists(ctx context.Context, key string) (bool, error) {
	// set client
	cl := r.read.Get()
	defer cl.Close()

	var ok bool
	var err error
	ok, err = redis.Bool(cl.Do("EXISTS", key))
	if err != nil {
		logger.LogRed(err.Error())
		return ok, err
	}

	return ok, nil
}

// Delete method
func (r *RedisStore) Delete(ctx context.Context, key string) error {
	// set client
	cl := r.write.Get()
	defer cl.Close()

	_, err := cl.Do("DEL", key)
	if err != nil {
		logger.LogRed(err.Error())
		return err
	}

	return nil
}
