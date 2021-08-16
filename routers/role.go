package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"jwtDemo/model"
	"net/http"
)


//获取所有的角色列表
func FindAllRole(c *gin.Context) {
	var roles  = GlobalService.FindAllRole()
	for k,v :=range roles  {
		roles[k].Children = GlobalService.GetAllPermByRoleId(v.RoleId)
	}
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.SUCCESS,
			"msg":   consts.GetMsg(consts.SUCCESS),
			"data":  roles,
			"count": GlobalService.GetRoleCount(),
		})
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
		//删除token
		GlobalService.DelToken(fmt.Sprintf("%s%d",model.TokenKey,user.UserId))

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
//根据roleid查询所有的信息
func GetMenuById(c *gin.Context){
	var role  model.Role
	if c.BindJSON(&role) == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  consts.GetMsg(consts.SUCCESS),
			"data": GlobalService.GetAllPermByRoleId(role.RoleId),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  consts.GetMsg(consts.SUCCESS),
			"data": nil,
		})
	}
}
//删除对应的menu和用户id
func DeleteMenuAndRoleId(c *gin.Context) {
	var roleMenu  model.RoleMenu
	//没有错误
	if c.BindJSON(&roleMenu) == nil {
		if err:=GlobalService.DeleteMenuAndRoleId(roleMenu);err!=nil{
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_DELETE_ERROR,
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
