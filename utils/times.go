package utils

import (
	"time"
)

func GetYmd() string {
	var t int64 = time.Now().Unix()
	var s string = time.Unix(t, 0).Format("2006-01-02")
	return s
}

func GetYmds() string {
	var t int64 = time.Now().Unix()
	var s string = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	return s
}