package configs

import (
	"fmt"
	"os"
)

// Environment additional in this service
type Environment struct {
	ServicePath string
}

var env Environment

// GetEnv get global additional environment
func GetEnv() *Environment {
	return &env
}

func loadAdditionalEnv() {
	env.ServicePath = lookupEnv("SERVICE_PATH")
}

func lookupEnv(envName string) string {
	env, ok := os.LookupEnv(envName)
	if !ok {
		fmt.Printf("missing %s environment will replace with empty value\n", envName)
	}

	return env
}
