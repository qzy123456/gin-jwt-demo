package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

)

type Config struct {
	Dsn          []string
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

	return engine
}
