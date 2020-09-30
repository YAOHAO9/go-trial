package logger

import (
	"log"
	"os"
	"time"

	"github.com/YAOHAO9/pine/application/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// LogType 当前日志类型
var LogType int

// SetLogMode 设置log模式
func SetLogMode(logType int) {
	LogType = logType

	logrus.AddHook(&errorHook{})
	logrus.SetReportCaller(true)

	if LogType == LogTypeEnum.File {
		// 产品模式
		logrus.SetLevel(logrus.ErrorLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})

		path, _ := os.Getwd()

		writer, err := rotatelogs.New(
			path+"/log/"+config.GetServerConfig().ID+"_%m_%d.log",
			rotatelogs.WithMaxAge(30*time.Second),
			rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
		)

		if err != nil {
			log.Fatalf("create file log.txt failed: %v", err)
		}

		logrus.SetOutput(writer)
	}

	if LogType == LogTypeEnum.Console {
		// 开发模式
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetFormatter(ErrorFormatter{})
	}
}
