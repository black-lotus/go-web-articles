package database

import (
	"context"
	"os"
	"testing"

	"webarticles/pkg/codebase/interfaces"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/spf13/cast"
)

func TestInitRedis(t *testing.T) {
	os.Setenv("REDIS_READ_HOST", gofakeit.Word())
	os.Setenv("REDIS_READ_PORT", cast.ToString(1))
	os.Setenv("REDIS_READ_AUTH", gofakeit.Word())
	os.Setenv("REDIS_READ_TLS", gofakeit.Word())

	os.Setenv("REDIS_WRITE_HOST", gofakeit.Word())
	os.Setenv("REDIS_WRITE_PORT", cast.ToString(1))
	os.Setenv("REDIS_WRITE_AUTH", gofakeit.Word())
	os.Setenv("REDIS_WRITE_TLS", gofakeit.Word())

	tests := map[string]struct {
		name string
		want interfaces.RedisPool
	}{
		"Test #1 positive init db redis connection": {},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() { recover() }()

			ctx := context.Background()
			redis := InitRedis()
			redis.ReadPool()
			redis.WritePool()
			redis.Disconnect(ctx)
		})
	}
}
