package orm

import (
	"fmt"
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

	//writer, err := rotatelogs.New(
	//	"./runtime/logs",
	//	rotatelogs.WithLinkName("mysql_log"),
	//	rotatelogs.WithMaxAge(time.Duration(86400 * 3)*time.Second),
	//	rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(writer)
	//
	//engine.SetLogger(logger)
	//是否打印sql语句
	if c.Debug{
		engine.ShowSQL(true)
	}

	return engine
}
