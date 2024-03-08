package logging

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

func SetLogger() {
	logger, ok := os.LookupEnv("LOGGER")
	if !ok {
		zap.L().Warn("LOGGER environment variable not set, using default logger")
	}

	// turn to uncapitalized string
	logger = strings.ToLower(logger)

	// set zap to json or console
	switch logger {
	case "json":
		logger, err := zap.NewProduction()
		if err != nil {
			zap.L().Error("error creating prod logger", zap.Error(err))
		}
		zap.ReplaceGlobals(logger)
	case "console":
		logger, err := zap.NewDevelopment()
		if err != nil {
			zap.L().Error("error creating dev logger", zap.Error(err))
		}
		zap.ReplaceGlobals(logger)
	default:
		zap.L().Warn("LOGGER environment variable not set to json or console, using default logger")
	}
}
