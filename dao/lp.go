package dao

import (
	"net"
)

// Ip2City 获取IP地理位置信息
func (d *Dao) Ip2City(ip2 string) (city string, err error) {
	ip := net.ParseIP(ip2)
	record, err := d.GeoIp.City(ip)
	if err != nil {
		return
	}
	city = record.City.Names["zh-CN"]
	return
}
