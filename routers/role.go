package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"jwtDemo/model"
	"net/http"
)


//获取所有的角色列表
func FindAllRole(c *gin.Context) {
	var pageInfo model.Page
	//没有错误
	if c.BindJSON(&pageInfo) == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.SUCCESS,
			"msg":   consts.GetMsg(consts.SUCCESS),
			"data":  GlobalService.FindAllRole(pageInfo),
			"count": GlobalService.GetRoleCount(pageInfo),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.INVALID_PARAMS,
			"msg":  consts.GetMsg(consts.INVALID_PARAMS),
		})
	}

}
//插入一条数据
func SaveRole(c *gin.Context) {
	var user  model.Role
	//没有错误
	if c.BindJSON(&user) == nil {
		//检测有无用户
		if GlobalService.CheckRoleByName(user.RoleName){
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_EXIST_USER,
				"msg":   consts.GetMsg(consts.ERROR_EXIST_USER),
			})
			return
		}
		if err:=GlobalService.SaveRole(user);err != nil{
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR,
				"msg":   err.Error(),
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
//根据用户id删除
func DeleteRoleById(c *gin.Context) {
	var user  model.UserRoleNew
	//没有错误
	if c.BindJSON(&user) == nil {
		if !GlobalService.DeleteRoleById(user.RoleId){
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_DELETE_ERROR,
				"msg":   consts.GetMsg(consts.ERROR_DELETE_ERROR),
			})
			return
		}

		//删除完角色，要删除角色的menu
		GlobalService.DelUserMenuByRoleId(user.RoleId)

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
//根据用户id修改
func UpdateroleById(c *gin.Context) {
	var user  model.Role
	//没有错误
	if c.BindJSON(&user) == nil {
		//检测用户名不能重复
		if GlobalService.CheckRoleByName(user.RoleName){
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_EXIST_USER,
				"msg":   consts.GetMsg(consts.ERROR_EXIST_USER),
			})
			return
		}
		//修改失败
		if !GlobalService.UpdateRoleById(user){
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_UPDATE_ERROR,
				"msg":   consts.GetMsg(consts.ERROR_UPDATE_ERROR),
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

