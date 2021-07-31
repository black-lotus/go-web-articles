package database

import (
	"context"
	"os"
	"testing"

	"webarticles/pkg/codebase/interfaces"

	"github.com/brianvoe/gofakeit/v5"
)

func TestInitSQL(t *testing.T) {
	os.Setenv("SQL_DATABASE_NAME", gofakeit.Word())
	os.Setenv("SQL_DB_USER", gofakeit.Word())
	os.Setenv("SQL_DB_PASSWORD", gofakeit.Word())
	os.Setenv("SQL_DRIVER_NAME", gofakeit.Word())

	tests := map[string]struct {
		name string
		want interfaces.SQLDatabase
	}{
		"Test #1 positive init db sql connection": {},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() { recover() }()

			ctx := context.Background()
			redis := InitSQLDatabase()
			redis.GetSQLDB()
			redis.Disconnect(ctx)
		})
	}
}
