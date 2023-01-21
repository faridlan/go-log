package test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello World")
}

func TestLevel(t *testing.T) {

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Error(err)
	}

	logger.SetOutput(file)

	logger.Info("hello world")
	logger.Warn("hello warn")
	logger.Error("hello error")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
	logger.Error("Hello Logger")
}

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "faridlan").Info("Hello Info")

	logger.WithField("username", "faridlan").WithField("name", "faridlan nulhakim").Info("Hello Info")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "faridlan",
		"name":     "nullhakim",
	}).Info("Hello World")
}
