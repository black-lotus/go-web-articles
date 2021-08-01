package configs

import (
	"context"

	"webarticles/pkg/codebase/factory/dependency"
	"webarticles/pkg/codebase/interfaces"
	"webarticles/pkg/config"
	"webarticles/pkg/config/database"
	"webarticles/pkg/validator"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// LoadConfigs load selected dependency configuration in this service
func LoadConfigs(baseCfg *config.Config) (deps dependency.Dependency) {

	loadAdditionalEnv()
	baseCfg.LoadFunc(func(ctx context.Context) []interfaces.Closer {
		// init mysql
		mysqlDeps := database.InitSQLDatabase()

		//init redis
		redisDeps := database.InitRedis()

		// inject all service dependencies
		deps = dependency.InitDependency(
			dependency.SetValidator(validator.NewValidator()),
			dependency.SetSQLDatabase(mysqlDeps),
			dependency.SetRedisPool(redisDeps),
			// ... add more dependencies
		)

		return []interfaces.Closer{mysqlDeps, redisDeps} // throw back to config for close connection when application shutdown
	})

	return deps
}
