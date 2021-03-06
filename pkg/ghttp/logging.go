package ghttp

import "github.com/sirupsen/logrus"

type Level uint32

const (
	PanicLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func getLogLevel(logLevel string) int {
	switch logLevel {
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warning":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	case "trace":
		return TraceLevel
	default:
		return InfoLevel
	}
}

func ConfigureLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.FullTimestamp = true
	logger.SetFormatter(formatter)
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.Level(getLogLevel(logLevel)))
	return logger
}
