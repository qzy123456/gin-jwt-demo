package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"jwtDemo/model"
	"net/http"
)


//获取所有的角色列表
func FindAllMenu(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.SUCCESS,
			"msg":   consts.GetMsg(consts.SUCCESS),
			"data":  GlobalService.GetAllPerms(),
		})

}
//插入一条数据
func SaveMenu(c *gin.Context) {
	var menu  model.Menu
	//没有错误
	if c.BindJSON(&menu) == nil {
		//检测有无用户
		if GlobalService.CheckMenuByName(menu){
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_EXIST_MENU_URL,
				"msg":   consts.GetMsg(consts.ERROR_EXIST_MENU_URL),
			})
			return
		}
		if err:=GlobalService.SaveMenu(menu);err != nil{
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
func DeleteMenuById(c *gin.Context) {
	var user  model.Menu
	//没有错误
	if c.BindJSON(&user) == nil {
		if err :=GlobalService.DeleteMenuById(user.MenuId);err != nil{
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
//根据用户id修改
func UpdateMenuById(c *gin.Context) {
	var menu  model.Menu
	//没有错误
	if c.BindJSON(&menu) == nil {
		//检测用户名不能重复
		if GlobalService.CheckMenuByName(menu){
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_EXIST_MENU_URL,
				"msg":   consts.GetMsg(consts.ERROR_EXIST_MENU_URL),
			})
			return
		}
		//修改失败
		if err := GlobalService.UpdateMenuById(menu);err != nil{
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR_UPDATE_ERROR,
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
//给角色分配id
func GiveMenu(c *gin.Context)  {
	var user  model.RoleMenuNew
	//没有错误
	if c.BindJSON(&user) == nil {
		//先删除角色的menu
		GlobalService.DelUserMenuByRoleId(user.RoleId)
		if err:=GlobalService.SaveRoleMenu(user);err != nil{
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

