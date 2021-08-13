package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"net/http"
)

func (m *Middleware) CheckMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var path = c.Request.URL.Path
		//不需要检测的
		claims := c.MustGet("claims").(*CustomClaims)
		menus := m.Service.GetAllUserMenus(claims.ID)
		fmt.Println(menus)
		var isSet = false
		for _, value := range menus {
			//在
			if value == path {
				isSet = true
			}
		}
		//不在权限列表，并且没有在
		if !isSet {
			if _, ok := m.NoCheckAction[path]; !ok {
				c.JSON(http.StatusOK, gin.H{
					"code": consts.FORBIDDEN,
					"msg":  consts.GetMsg(consts.FORBIDDEN),
				})
				c.Abort()
				return
			}
		}
		// 处理请求
		c.Next()
	}
}
