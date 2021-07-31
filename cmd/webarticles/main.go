package main

import (
	"fmt"
	"runtime/debug"

	service "webarticles/internal"
	"webarticles/pkg/codebase/app"
	"webarticles/pkg/config"
)

const (
	serviceName = "webarticles"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\x1b[31;1mFailed to start %s service: %v\x1b[0m\n", serviceName, r)
			fmt.Printf("Stack trace: \n%s\n", debug.Stack())
		}
	}()

	cfg := config.Init(fmt.Sprintf("cmd/%s/", serviceName))
	defer cfg.Exit()

	srv := service.NewService(serviceName, cfg)
	app.New(srv).Run()
}
