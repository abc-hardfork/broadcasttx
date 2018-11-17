package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLog() error {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05", DisableTimestamp: false})

	// set log level
	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		return err
	}

	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetOutput(file)

	return nil
}
