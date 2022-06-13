package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// Interface
type Interface interface {
	Debug(args ...interface{})
	// TO DO DebugWithFields(msg interface{}, f Fields)
	Info(args ...interface{})
}

// Logger is a wrapper around logrus.Logger
type Logger struct {
	logger *logrus.Logger
}

// Fields is a wrapper around logrus.Fields
type Fields logrus.Fields

var _ Interface = (*Logger)(nil)

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
	//	l.logger.WithFields(logrus.Fields{"host": os.Getenv("HOSTNAME")}).Debug(args)
}

/* TO-DO
func (l *Logger) DebugWithFields(msg interface{}, f Fields) {
	l.logger.WithFields(logrus.Fields(f)).Debug(msg)
}
*/

// Info sends an unstructured informational log message.
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args)
}
