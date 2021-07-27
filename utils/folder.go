package utils

import (
	"os"
	"path/filepath"
	"time"
)

//定义一个创建文件目录的方法
func Mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006")
	folderPath := filepath.Join(basePath,folderName)
	//使用mkdirall会创建多层级目录
	os.MkdirAll(folderPath, os.ModePerm)
	return  folderPath
}