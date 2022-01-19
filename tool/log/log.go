package log

import (
	"io/ioutil"
	sys "log/syslog"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/syslog"
)

var (
	Logger *logrus.Entry
)

const ENV_PROD = "prod"

func init() {
	// 默认日志输出
	InitLoggerJson("api-default")
}

func InitLoggerJson(tag string) {
	logAddr := ""
	runEnv := GetRunMode()
	logrus.SetReportCaller(true) // 输出文件名，行号和函数名
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:03:04",
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := path.Base(f.File)
			return f.Function, fileName
		},
	})
	Logger = logrus.WithFields(logrus.Fields{
		"env": runEnv,
	})
	if runEnv == ENV_PROD {
		// 往日志管理平台输送日志
		syslog.NewSyslogHook("udp", logAddr, sys.LOG_INFO, tag)
		// 关闭标准日志输出
		logrus.SetOutput(ioutil.Discard)
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func InitLoggerText() {
	runEnv := GetRunMode()
	logrus.SetReportCaller(true) // 输出文件名，行号和函数名
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:03:04",
	})
	Logger = logrus.WithField("env", runEnv)
	if runEnv == ENV_PROD {
		// 关闭标准日志输出
		logrus.SetOutput(ioutil.Discard)
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func GetRunMode() string {
	return os.Getenv("RUN_MODE")
}
