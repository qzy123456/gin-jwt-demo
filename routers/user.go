package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"jwtDemo/servcie"
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
	c.JSON(http.StatusOK, gin.H{
		"code": consts.SUCCESS,
		"msg":    consts.GetMsg(consts.SUCCESS),
		//"data":servcie.FindMenuById(1),
		"data2":servcie.GetAllPerm2(1),
		//"data3":servcie.FindMenus(1),
		"data4":servcie.GetAllPerm4(0,servcie.FindMenus(1)),
	})
}



