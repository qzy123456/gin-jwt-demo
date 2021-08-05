package orm

import (
	"fmt"
	"github.com/arthurkiller/rollingwriter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Config struct {
	Dsn          []string
	Debug        bool
	MaxId  		 int
	MaxOpen      int
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (engine *xorm.EngineGroup) {
	//连接数据库
	engine, err := xorm.NewEngineGroup("mysql", c.Dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	//连接测试
	if err := engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	//最大空闲连接数
	engine.SetMaxIdleConns(c.MaxId)
	//最大打开连接数
	engine.SetMaxOpenConns(c.MaxOpen)
	config := rollingwriter.Config{
		LogPath:       "./runtime/logs",       //日志路径
		TimeTagFormat: "060102150405", //时间格式串
		FileName:      "mysql_log",   //日志文件名
		MaxRemain:     3,              //配置日志最大存留数
		RollingPolicy:      rollingwriter.VolumeRolling, //配置滚动策略 norolling timerolling volumerolling
		RollingTimePattern: "* * * * * *",               //配置时间滚动策略
		RollingVolumeSize:  "1M",                        //配置截断文件下限大小
		WriterMode: "none",
		BufferWriterThershould: 256,
		Compress: true,
	}

	writer, err := rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		panic(err)
	}

	var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(writer)

	engine.SetLogger(logger)
	//是否打印sql语句
	if c.Debug{
		engine.ShowSQL(true)
	}

	return engine
}
