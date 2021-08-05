package model

import "github.com/dgrijalva/jwt-go"

// User 用户类
type User struct {
	UserId         int `json:"userId" xorm:"pk user_id autoincr "`
	Username       string `json:"userName" xorm:"unique"`
	Password     string `json:"password" xorm:"column:password"`
	Enabled      int `json:"enabled" xorm:"column:enabled"`
	CreateTime   string `json:"create" xorm:"column:create_time"`
	LastTime     string `json:"last" xorm:"column:last_time"`
}

func (User) TableName()string  {
	return "tbl_user"
}
// LoginReq 登录请求参数类
type LoginReq struct {
	Username string `json:"username"`
	Password   string `json:"password"`
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID    int `json:"userId"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

