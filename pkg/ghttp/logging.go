package ghttp

import "github.com/sirupsen/logrus"

type Level uint32

const (
	PanicLevel = iota
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
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	customFormatter.FullTimestamp = true
	logger.SetFormatter(customFormatter)
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.Level(getLogLevel(logLevel)))
	return logger
}
