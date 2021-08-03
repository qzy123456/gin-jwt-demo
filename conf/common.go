package conf

import "errors"

var (
	PageSize          uint    = 10
	Version           string  = "0.3.9"
	Upload            string  = "upload/"
	Dir               string  = "configs/"
	MysqlConf         string  = Dir + "mysql.json"
	IsCompireTemplate bool    = true //是否编译静态模板到二进制
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
)
