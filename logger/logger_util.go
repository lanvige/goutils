package logger

import (
	"github.com/sirupsen/logrus"
)

var xlog = logrus.New()

// Debug send log debug to logstash
func Debug(args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Debug(args...)
}

// Debugf send log debug to logstash
func Debugf(format string, args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Debugf(format, args...)
}

// Info send log info to logstash
func Info(args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Info(args...)
}

// Infof send log info to logstash
func Infof(format string, args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Infof(format, args...)
}

// Warn send log warn to logstash
func Warn(args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Warn(args...)
}

// Warnf send log warn to logstash
func Warnf(format string, args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Warnf(format, args...)
}

// Error send log error to logstash
func Error(args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Error(args...)
}

// Errorf send log error to logstash
func Errorf(format string, args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Errorf(format, args...)
}

// Fatal send log fatal to logstash
func Fatal(args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Fatal(args...)
}

// Fatalf send log fatal to logstash
func Fatalf(format string, args ...interface{}) {
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.SetLevel(logrus.DebugLevel)

	xlog.Fatalf(format, args...)
}
