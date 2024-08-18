package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func init() {
	// create a new logger instance
	log = logrus.New()
	// set output to stdout
	log.SetOutput(os.Stdout)
	// set log formatter
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})
}

func Info(message string, fields logrus.Fields) {
	log.WithFields(fields).Info(message)
}

func InfoF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Infof(format, args...)
}

func Debug(message string, fields logrus.Fields) {
	log.WithFields(fields).Debug(message)
}

func DebugF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Debugf(format, args...)
}

func Error(message string, fields logrus.Fields) {
	log.WithFields(fields).Error(message)
}

func ErrorF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Errorf(format, args...)
}

func Fatal(message string, fields logrus.Fields) {
	log.WithFields(fields).Fatal(message)
}

func FatalF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Fatalf(format, args...)
}

func Panic(message string, fields logrus.Fields) {
	log.WithFields(fields).Panic(message)
}

func PanicF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Panicf(format, args...)
}

func Warn(message string, fields logrus.Fields) {
	log.WithFields(fields).Warn(message)
}

func WarnF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Warnf(format, args...)
}
