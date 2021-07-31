package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// Env model
type Env struct {
	RootApp string

	// UseREST env
	UseREST bool
	// UseSwagger env
	UseSwagger bool

	// JSONSchemaDir env
	JSONSchemaDir string

	// Development env checking, this env for debug purpose
	Development string

	// Env on application
	Environment string

	// RESTPort config
	RESTPort uint16

	DebugMode bool
}

func loadBaseEnv(serviceLocation string, targetEnv *Env) {

	// load main .env and additional .env in app
	if err := godotenv.Load(serviceLocation + ".env"); err != nil {
		fmt.Println("System cannot load env file in project. System will try read config from env in os variable ...")
	}

	rootApp, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Setenv("APP_PATH", rootApp)
	env.RootApp = rootApp

	// ------------------------------------

	useREST, ok := os.LookupEnv("USE_REST")
	if !ok {
		fmt.Println("missing USE_REST environment will replace with the default `false`")
		useREST = "false"
	}
	env.UseREST, _ = strconv.ParseBool(useREST)

	useSWAGGER, ok := os.LookupEnv("USE_SWAGGER")
	if !ok {
		fmt.Println("missing USE_SWAGGER environment will replace with the default `false`")
		useSWAGGER = "false"
	}
	env.UseSwagger, _ = strconv.ParseBool(useSWAGGER)

	// ------------------------------------

	if env.UseREST {
		if restPort, ok := os.LookupEnv("REST_HTTP_PORT"); !ok {
			panic("missing REST_HTTP_PORT environment")
		} else {
			port, err := strconv.Atoi(restPort)
			if err != nil {
				panic("REST_HTTP_PORT environment must in integer format")
			}
			env.RESTPort = uint16(port)
		}
	}

	// ------------------------------------
	env.Environment = os.Getenv("ENVIRONMENT")
	env.JSONSchemaDir, ok = os.LookupEnv("JSON_SCHEMA_DIR")
	if !ok {
		panic("missing JSON_SCHEMA_DIR environment")
	}
}
