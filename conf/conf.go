package conf

import (
	"flag"
	"fmt"
	"jwtDemo/library/database/orm"
	"path"
	"github.com/spf13/viper"
	cache "jwtDemo/library/cache/redis"
	"jwtDemo/library/xlog"
)

const (
	_localConfPath    = "configs/dev"
	_localConf        = "m3_admin.json"
)

var (
	confPath string
	Conf     Config
)

type App struct {
	GlobalCachePrefix string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int64
	WriteTimeout int64
	Address      string
}

type Redis struct {
	Global      *cache.Config
	UserCluster *cache.Config
}



type Config struct {
	// m3_admin.json
	App      *App
	Server   *Server
	Log      *xlog.Config
	Redis    *Redis
	Db       *orm.Config
}

//读取命令行启动参数，没有就要用dev
func init() {
	flag.StringVar(&confPath, "conf", _localConfPath, "configs file path.")
}

// Init Config
func Init() (err error) {
	fmt.Println("==== Init m3 from " + confPath + " ====")

	// 读取 server 配置
	confViper := viper.New()
	confViper.SetConfigFile(path.Join(confPath, _localConf))
	err = confViper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error %s configs file: %s", path.Join(confPath, _localConf), err)
	}
	err = confViper.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	fmt.Println(Conf.Db.Dsn)
	Conf.Log.LevelMode = Conf.Server.RunMode
	return
}
