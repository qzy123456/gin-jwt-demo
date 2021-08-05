package xlog

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func newLfsHook(c *Config, suffix string) *lfshook.LfsHook {
	var bathLogPath, bathLinkPath, bathErrorPath string
	if suffix != "" {
		bathLogPath = c.Path + suffix + "." + c.FilePrefix
		bathLinkPath = c.Path + suffix + ".log"
		bathErrorPath = c.Path + suffix + "_error" + ".log"
	} else {
		bathLogPath = c.Path + c.FilePrefix
		bathLinkPath = c.Path + "go.log"
		bathErrorPath = c.Path + "error.log"
	}
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	 `WithMaxAge` 设置文件清理前的最长保存时间
	 `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。
	writer, err := rotatelogs.New(
		bathLogPath,
		rotatelogs.WithLinkName(bathLinkPath),
		//rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		//rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(c.RotationTime)*time.Second),
	)

	if err != nil {
		logrus.Errorf("configs local file system logger error. %+v", err)
	}

	errWriter, err := rotatelogs.New(
		bathErrorPath,
		//rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(c.RotationTime)*time.Second),
	)

	if err != nil {
		logrus.Errorf("configs local file system logger error. %+v", err)
	}

	return lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: errWriter,
			logrus.FatalLevel: errWriter,
			logrus.PanicLevel: errWriter,
		}, &JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FieldMap: FieldMap{
				logrus.FieldKeyLevel: "logLevel",
				logrus.FieldKeyMsg:   "msg",
				logrus.FieldKeyTime:  "@timestamp",
			},
		})
}
