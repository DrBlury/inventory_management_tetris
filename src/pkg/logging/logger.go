package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetLogger() {
	logger, ok := os.LookupEnv("LOGGER")
	if !ok {
		log.Fatalln("ENV ERROR: 'LOGGER' is missing")
	}

	if logger == "JSON" || logger == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
	}
}
