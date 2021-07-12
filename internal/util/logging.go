package util

import (
	"github.com/sirupsen/logrus"
	"os"
)

func GetLogLevelFromEnv() logrus.Level {
	logLevel, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL is not set, let's default to debug
	if !ok {
		logLevel = "DEBUG"
	}

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		lvl = logrus.DebugLevel
	}

	return lvl
}
