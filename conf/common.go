package conf

var (
	PageSize          uint    = 10
	Version           string  = "0.3.9"
	Upload            string  = "upload/"
	Dir               string  = "configs/"
	MysqlConf         string  = Dir + "mysql.json"
	IsCompireTemplate bool    = true //是否编译静态模板到二进制
)
