package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// Interface
type Interface interface {
	Debug(args ...interface{})
	//	DebugWithFields(fields interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

// Logger is a wrapper around logrus.Logger
type Logger struct {
	logger *logrus.Logger
}

// Fields is a wrapper around logrus.Fields
type Fields logrus.Fields

var _ Interface = (*Logger)(nil)

// New creates a new logger.
func New(loglevel string) *Logger {
	var l logrus.Level

	switch strings.ToLower(loglevel) {
	case "debug":
		l = logrus.DebugLevel
	case "info":
		l = logrus.InfoLevel
	case "warn":
		l = logrus.WarnLevel
	case "error":
		l = logrus.ErrorLevel
	case "fatal":
		l = logrus.FatalLevel
	case "panic":
		l = logrus.PanicLevel
	default:
		l = logrus.InfoLevel
	}

	//	logrus.SetLevel(l)
	logger := logrus.New()
	logger.SetLevel(l)
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{logger: logger}
}

// Debug sends an unstructured debug log message.
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

/*
func (l *Logger) DebugWithFields(fields Fields) {
	//	l.logger.WithFields(logrus.Fields(f)).Debug(msg)
	//	l.logger.WithFields(logrus.Fields{"host": os.Getenv("HOSTNAME")}).Debug(args)
	l.logger.WithFields(logrus.Fields(fields))
}
*/

// Info sends an unstructured informational log message.
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args)
}

// Warn sends an unstructured warning log message.
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args)
}

// Error sends an unstructured error log message.
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args)
}

// Fatal sends an unstructured fatal log message.
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args)
}

// Panic sends an unstructured panic log message.
func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args)
}
