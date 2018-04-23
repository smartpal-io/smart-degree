package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func NewLogger(loggerName string) *log.Entry{
	log.SetFormatter(&log.TextFormatter{})
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	return log.WithField("logger", loggerName)
}
