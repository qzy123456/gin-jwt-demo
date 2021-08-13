package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"jwtDemo/model"
	"net/http"
)

func GetLogs(c *gin.Context)  {
	var pageInfo model.Page
	//没有错误
	if c.BindJSON(&pageInfo) == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.SUCCESS,
			"msg":   consts.GetMsg(consts.SUCCESS),
			"data":  GlobalService.FindAllLog(pageInfo),
			"count": GlobalService.GetLogCount(pageInfo),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.INVALID_PARAMS,
			"msg":  consts.GetMsg(consts.INVALID_PARAMS),
			"data" :c.BindJSON(&pageInfo).Error(),
		})
	}
}
