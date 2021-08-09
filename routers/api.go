package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	myjwt "jwtDemo/middleware"
	"jwtDemo/model"
	"net/http"
)

// Login 登录
func Login(c *gin.Context) {
	var loginReq model.LoginReq
	if c.BindJSON(&loginReq) == nil {
		user := GlobalService.CheckLogin(loginReq)
		fmt.Println(user)
		if user != nil {
			GlobalService.GenerateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": consts.INVALID_PARAMS,
				"msg":  consts.GetMsg(consts.INVALID_PARAMS),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.INVALID_PARAMS,
			"msg":  consts.GetMsg(consts.INVALID_PARAMS),
		})
	}
}



// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

// 刷新token
func  Refresh(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "请求未携带token，无权限访问",
		})
		c.Abort()
		return
	}
	//刷新token
	tokens, err := GlobalService.RefreshToken(token)
	if err != nil {
		fmt.Println("验证失败",err)
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"data":   tokens,
	})
}

