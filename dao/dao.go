package dao

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/go-xorm/xorm"
	"github.com/oschwald/geoip2-golang"
	"jwtDemo/conf"
	cache "jwtDemo/library/cache/redis"
	"jwtDemo/library/database/orm"
	"log"
)

type Dao struct {
	C           *conf.Config
	Db          *xorm.EngineGroup
	GlobalCache *cache.Pool
	UsersCache  *cache.Pool
	HttpClient  *resty.Client
	GeoIp       *geoip2.Reader
	Jwt         *conf.Jwt
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		C:           c,
		Db:          orm.NewMySQL(c.Db),
		GlobalCache: cache.NewPool(c.Redis.Global),
		UsersCache:  cache.NewPool(c.Redis.UserCluster),
		HttpClient:  resty.New(),
	}
	return
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.pingRedis(ctx); err != nil {
		return
	}
	return
}

func (d *Dao) pingRedis(ctx context.Context) (err error) {
	conn, err := d.GlobalCache.Conn.GetContext(ctx)
	if err != nil {
		return
	}
	defer conn.Close()
	if _, err = conn.Do("SETEX", "ping",300, "pong"); err != nil {
		log.Fatalf("conn.Set(PING) error(%v)", err)
	}
	return
}

func (d *Dao) Close() (err error) {
	err = d.GeoIp.Close()
	return
}
