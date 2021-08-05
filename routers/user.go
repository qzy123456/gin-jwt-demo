package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"jwtDemo/middleware"
	"jwtDemo/model"
	"net/http"
)

//获取所有的后台登陆用户
func FindAllUser(c *gin.Context) {
	var pageInfo model.Page
	//没有错误
	if c.BindJSON(&pageInfo) == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.SUCCESS,
			"msg":   consts.GetMsg(consts.SUCCESS),
			"data":  GlobalService.FindAllUser(pageInfo),
			"count": GlobalService.GetUserCount(pageInfo),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.INVALID_PARAMS,
			"msg":  consts.GetMsg(consts.INVALID_PARAMS),
		})
	}

}

//根据用户id，获取所有的菜单列表，封装成tree
func FindMenuById(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  consts.GetMsg(consts.SUCCESS),
			"data": GlobalService.GetAllPerm2(claims.ID),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  consts.GetMsg(consts.SUCCESS),
			"data": nil,
		})
	}

}

//插入一条数据
func SaveUser(c *gin.Context) {
	var user  model.User
	//没有错误
	if c.BindJSON(&user) == nil {
		//检测有无用户
		if GlobalService.CheckUserByName(user.Username){
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_EXIST_USER,
				"msg":   consts.GetMsg(consts.ERROR_EXIST_USER),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.SUCCESS,
			"msg":   consts.GetMsg(consts.SUCCESS),
			"data":  GlobalService.SaveUser(user),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.INVALID_PARAMS,
			"msg":  consts.GetMsg(consts.INVALID_PARAMS),
		})
	}
}
//根据用户id删除
func DeleteById(c *gin.Context) {
	var user  model.User
	//没有错误
	if c.BindJSON(&user) == nil {
		if !GlobalService.DeleteById(5){
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_DELETE_ERROR,
				"msg":   consts.GetMsg(consts.ERROR_DELETE_ERROR),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.SUCCESS,
			"msg":   consts.GetMsg(consts.SUCCESS),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.INVALID_PARAMS,
			"msg":  consts.GetMsg(consts.INVALID_PARAMS),
		})
	}
}
