package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func SetLogger(level string) {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Printf("error set level logrus\n")
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrusLevel)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetOutput(os.Stdout)
}
