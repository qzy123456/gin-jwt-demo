package model

import "github.com/dgrijalva/jwt-go"

const (
	dbName     = "myBlog.db"
	userBucket = "user"
)

// User 用户类
type User struct {
	UserId         int `json:"userId" gorm:"primary_key;column:user_id"`
	Username       string `json:"userName" gorm:"column:username"`
	Password     string `json:"password" gorm:"column:password"`
	Enabled      int `json:"enabled" gorm:"column:enabled"`
	CreateTime   string `json:"create" gorm:"column:create_time"`
	LastTime     string `json:"last" gorm:"column:last_time"`
}

func (User) TableName()string  {
	return "tbl_user"
}
// LoginReq 登录请求参数类
type LoginReq struct {
	Phone string `json:"mobile"`
	Pwd   string `json:"pwd"`
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID    int `json:"userId"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

// LoginCheck 登录验证
func LoginCheck(loginReq LoginReq) (bool, User, error) {

	resultUser := User{}
	resultBool := true

	loginReq.Phone = resultUser.Username
	loginReq.Pwd = resultUser.Password

	return resultBool, resultUser, nil
}



