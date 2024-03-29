package logging

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

func SetLogger() *zap.SugaredLogger {
	loggerVar, ok := os.LookupEnv("LOGGER")
	loggerVar = strings.ToLower(loggerVar)

	// set zap to json or console
	if ok && loggerVar == "json" {
		logger, err := zap.NewProduction()
		if err != nil {
			zap.L().Sugar().Warnf("error creating prod logger: ", err)
		}
		zap.ReplaceGlobals(logger)
		return logger.Sugar()
	}

	// default to console
	zap.L().Warn("LOGGER environment variable not set to json, using default logger")
	logger, err := zap.NewDevelopment()
	if err != nil {
		zap.L().Sugar().Warnf("error creating dev logger: ", err)
	}
	zap.ReplaceGlobals(logger)
	return logger.Sugar()
}
