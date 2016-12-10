package log

import (
	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

// Debugf logs a message in sprintf format at debug level.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof logs a message in sprintf format at info level.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warningf logs a message in sprintf format at warning level.
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

// Errorf logs a message in sprintf format at error level.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf logs a message in sprintf format at fatal level.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Warning logs an error object at warning level.
func Warning(err error) {
	log.Warning(err.Error())
}

// Error logs an error object at error level.
func Error(err error) {
	log.Error(err)
}

// Fatal logs an error object at fatal level.
func Fatal(err error) {
	log.Fatal(err)
}
