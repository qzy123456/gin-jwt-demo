package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/servcie"
	"net/http"
)


func FindAllUser(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "注册成功！",
		"data":GlobalService.FindAllUser(),
	})
}

func FindMenuById(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "注册成功！",
		//"data":servcie.FindMenuById(1),
		"data2":servcie.GetAllPerm2(1),
		//"data3":servcie.FindMenus(1),
		"data4":servcie.GetAllPerm4(0,servcie.FindMenus(1)),
	})
}


