package log

import (
	"io/ioutil"
	sys "log/syslog"
	"os"

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
	runEnv := os.Getenv("run_env")
	logrus.SetFormatter(&logrus.JSONFormatter{})
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
	runEnv := os.Getenv("run_env")
	logrus.SetFormatter(&logrus.TextFormatter{})
	Logger = logrus.WithField("env", runEnv)
	if runEnv == ENV_PROD {
		// 关闭标准日志输出
		logrus.SetOutput(ioutil.Discard)
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
