package main

import (
	"fmt"
	"linuxcode/inventory_manager/pkg/app"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

// Application metadata that is set at compile time.
// nolint
var (
	version     string
	buildDate   string
	description = "application"
	commitHash  string
	commitDate  string
)

// main just loads config and inits logger. Rest is done in app.Run.
func main() {
	appCfg, err := app.LoadConfig(
		version,
		buildDate,
		description,
		commitHash,
		commitDate,
	)
	if err != nil {
		fmt.Printf("could not load config: %s", err.Error())
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	err = app.Run(
		appCfg,
		quit,
	)

	if err != nil {
		zap.L().Error("error running app", zap.Error(err))
	}
}
