package redis

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Pool struct {
	c    *Config
	Conn *redis.Pool
}

// Config client settings.
// MaxIdle 池中最大空闲连接数。
// MaxActive 给定时间池分配的最大连接数。 如果为零，则池中的连接数没有限制。
// IdleTimeout 在此期间保持空闲后关闭连接。 如果该值为零，则不关闭空闲连接。 应用程序应将超时设置为小于服务器超时的值。
type Config struct {
	Host        string
	Password    string
	Db          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int64
}

// NewPool creates a new pool.
func NewPool(conf *Config) *Pool {
	return &Pool{
		c: conf,
		Conn: &redis.Pool{
			MaxIdle:     conf.MaxIdle,
			MaxActive:   conf.MaxActive,
			IdleTimeout: time.Duration(conf.IdleTimeout) * time.Second,
			Dial: func() (redis.Conn, error) {
				// 拨号是应用程序提供的用于创建和配置连接的功能。
				c, err := redis.Dial("tcp", conf.Host)
				if err != nil {
					log.Fatalf("redis.Setup, fail to dial: %v", err)
					return nil, err
				}
				if conf.Password != "" {
					if _, err := c.Do("AUTH", conf.Password); err != nil {
						c.Close()
						log.Fatalf("redis.Setup, fail to auth: %v", err)
						return nil, err
					}
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				if err != nil {
					log.Fatalf("redis.Setup, fail to Conn: %v", err)
				}
				return nil
			},
		},
	}
}

func (p *Pool) Close() error {
	return p.Conn.Close()
}
