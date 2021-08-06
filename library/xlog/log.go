package xlog

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Path         string
	FilePrefix   string
	LevelMode    string
	RotationTime int64
}

// 初始化 logger 服务
func Init(c *Config, suffix string) *logrus.Logger {
	logger := logrus.New()
	logger.AddHook(newLfsHook(c, suffix))
	
	/*
	   如果日志级别不是debug就不要打印日志到控制台了
	*/
	switch c.LevelMode {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
		logger.SetOutput(os.Stderr)
	case "info":
		setNull(logger)
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		setNull(logger)
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		setNull(logger)
		logger.SetLevel(logrus.ErrorLevel)
		logger.SetReportCaller(true)
	default:
		setNull(logger)
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}

// 取消标准输出
func setNull(logger *logrus.Logger) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logger.Errorf("err: %+v", err)
	}
	writer := bufio.NewWriter(src)
	logger.SetOutput(writer)
}
