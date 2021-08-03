package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"jwtDemo/middleware"
	"net/http"
)

//获取所有的后台登陆用户
func FindAllUser(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code": consts.SUCCESS,
		"msg":    consts.GetMsg(consts.SUCCESS),
		"data":GlobalService.FindAllUser(),
	})
}

//根据用户id，获取所有的菜单列表，封装成tree
func FindMenuById(c *gin.Context)  {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	if claims != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":    consts.GetMsg(consts.SUCCESS),
			"data":GlobalService.GetAllPerm2(claims.ID),
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  consts.GetMsg(consts.SUCCESS),
			"data": nil,
		})
	}

}



