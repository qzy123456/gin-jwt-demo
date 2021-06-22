package initialize

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {

     return GormMysql()

}

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db_rbac?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	//defer db.Close()
	db.LogMode(true)
	return  db
}
