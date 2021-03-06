package dependency

import (
	"webarticles/pkg/codebase/interfaces"
)

// Dependency base
type Dependency interface {
	GetSQLDatabase() interfaces.SQLDatabase
	GetRedisPool() interfaces.RedisPool
	GetValidator() interfaces.Validator
}

// Option func type
type Option func(*deps)

type deps struct {
	sqlDB     interfaces.SQLDatabase
	redisPool interfaces.RedisPool
	validator interfaces.Validator
}

// SetSQLDatabase option func
func SetSQLDatabase(db interfaces.SQLDatabase) Option {
	return func(d *deps) {
		d.sqlDB = db
	}
}

// SetRedisPool option func
func SetRedisPool(db interfaces.RedisPool) Option {
	return func(d *deps) {
		d.redisPool = db
	}
}

// SetValidator option func
func SetValidator(validator interfaces.Validator) Option {
	return func(d *deps) {
		d.validator = validator
	}
}

// InitDependency constructor
func InitDependency(opts ...Option) Dependency {
	opt := new(deps)

	for _, o := range opts {
		o(opt)
	}

	return opt
}

func (d *deps) GetSQLDatabase() interfaces.SQLDatabase {
	return d.sqlDB
}

func (d *deps) GetRedisPool() interfaces.RedisPool {
	return d.redisPool
}

func (d *deps) GetValidator() interfaces.Validator {
	return d.validator
}
