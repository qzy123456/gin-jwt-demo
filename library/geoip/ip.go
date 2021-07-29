package geoip

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
)

type Config struct {
	Path string
}

func New(c *Config) *geoip2.Reader {
	db, err := geoip2.Open(c.Path)
	if err != nil {
		panic(fmt.Sprintf("failed to connect GeoLite2-City.mmdb,err: %v", err))
	}
	return db
}
