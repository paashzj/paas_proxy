package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "logs/paas-proxy.log",
		MaxSize:    50,
		MaxAge:     7,
		MaxBackups: 200,
		LocalTime:  false,
		Compress:   false,
	})
}
