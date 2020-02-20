// Package glogger represents a generic logging interface for go

package glogger

import (
	"github.com/pkg/errors"
)

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger

// Logger represent common interface for logging function
type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

func InitLogger(lc LogConfig) (Logger, error) {
	loggerType := lc.Code
	l, err := GetLogFactoryBuilder(loggerType).Build(&lc)
	if err != nil {
		return l, errors.Wrap(err, "")
	}
	return l, nil
}

