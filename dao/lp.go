package dao

import (
	"net"
)

// Ip2City 获取IP地理位置信息
func (d *Dao) Ip2City(ip2 string) (city string, err error) {
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ip2)
	record, err := d.GeoIp.City(ip)
	if err != nil {
		return
	}
	city = record.City.Names["zh-CN"]
	/* if len(record.Subdivisions) > 0 {
		fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["zh-CN"])
		city = record.Subdivisions[0].Names["zh-CN"] + "," + city
	} */
	return
	/* fmt.Printf("Russian country name: %v\n", record.Country.Names["zh-CN"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone) */
}
