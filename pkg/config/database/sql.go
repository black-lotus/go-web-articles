package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"webarticles/pkg/codebase/interfaces"
	"webarticles/pkg/logger"
)

type sqlInstance struct {
	db *sql.DB
}

func (s *sqlInstance) GetSQLDB() *sql.DB {
	return s.db
}

func (s *sqlInstance) Disconnect(ctx context.Context) (err error) {
	deferFunc := logger.LogWithDefer("sql: disconnect...")
	defer deferFunc()

	return s.db.Close()
}

// InitSQLDatabase return sql db read & write instance
func InitSQLDatabase() interfaces.SQLDatabase {
	deferFunc := logger.LogWithDefer("Load SQL connection...")
	defer deferFunc()

	inst := new(sqlInstance)

	dbName, ok := os.LookupEnv("SQL_DATABASE_NAME")
	if !ok {
		panic("missing SQL_DATABASE_NAME environment")
	}

	var err error
	descriptor := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=false&parseTime=true",
		os.Getenv("SQL_DB_USER"), os.Getenv("SQL_DB_PASSWORD"), os.Getenv("SQL_DB_HOST"), dbName)
	inst.db, err = sql.Open(os.Getenv("SQL_DRIVER_NAME"), descriptor)
	if err != nil {
		panic("SQL Read: " + err.Error())
	}

	return inst
}
