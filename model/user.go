package model

import "github.com/dgrijalva/jwt-go"

// User 用户类
type User struct {
	UserId         int `json:"userId" xorm:"pk user_id autoincr "`
	Username       string `json:"userName" xorm:"unique"`
	Password     string `json:"password" xorm:"-"`
	Enabled      bool `json:"enabled" xorm:"enabled"`
	CreateTime   string `json:"create" xorm:"create_time"`
	LastTime     string `json:"last" xorm:"last_time"`
	RoleName     string  `json:"role_name" xorm:"role_name"`
	Role        Role ` xorm:"extends"`
	UserRole    UserRoleNew ` xorm:"extends"`
}
// User 用户类
type UserNew struct {
	UserId         int `json:"userId" xorm:"pk user_id autoincr "`
	Username       string `json:"userName" xorm:"unique"`
	Password     string `json:"password" xorm:"password"`
	Enabled      int `json:"enabled" xorm:"enabled"`
	CreateTime   string `json:"create" xorm:"create_time"`
	LastTime     string `json:"last" xorm:"last_time"`
}
func (User) TableName()string  {
	return "tbl_user"
}
func (UserNew) TableName()string  {
	return "tbl_user"
}
// LoginReq 登录请求参数类
type LoginReq struct {
	Username string `json:"username"`
	Password   string `json:"password"`
}
// UpdatePass 更改密码的请求体
type UpdatePass struct {
	UserName  string `json:"userName"`
	Password   string `json:"password"`
	NewPassword   string `json:"new_password"`
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID    int `json:"userId"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

