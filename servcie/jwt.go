package servcie

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jwtDemo/conf"
	"jwtDemo/model"
	"log"
	"net/http"
	"time"
)

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	model.User
}

// 生成令牌
func (s *Service) GenerateToken(c *gin.Context, user *model.User) {

	claims := model.CustomClaims{
		user.UserId,
		user.Username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 864000), // 过期时间 一天
			Issuer:   s.Jwt,                   //签名的发行者
		},
	}

 	token, err := s.CreateToken(claims)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  *user,
		Token: token,
	}
	//登陆成功，返回token，用户数据
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

// CreateToken 生成一个token
func  (s *Service) CreateToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.Jwt))
}
// 更新token
func (s *Service) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Jwt), nil
	})
	fmt.Println("解析失败",err)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", conf.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return "", conf.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", conf.TokenNotValidYet
			} else {
				return "", conf.TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return s.CreateToken(*claims)
	}
	return "", conf.TokenInvalid
}