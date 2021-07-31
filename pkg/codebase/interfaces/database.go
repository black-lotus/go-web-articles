package interfaces

import (
	"database/sql"

	"github.com/gomodule/redigo/redis"
)

// SQLDatabase abstraction
type SQLDatabase interface {
	GetSQLDB() *sql.DB
	Closer
}

// RedisPool abstraction
type RedisPool interface {
	ReadPool() *redis.Pool
	WritePool() *redis.Pool
	Store() Store
	Closer
}
