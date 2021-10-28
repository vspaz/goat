package ghttp

import "github.com/sirupsen/logrus"

func ConfigureLogger() *logrus.Logger{
	logger := logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	customFormatter.FullTimestamp = true
	logger.SetFormatter(customFormatter)
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
