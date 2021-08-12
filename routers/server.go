package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"jwtDemo/consts"
	"jwtDemo/utils"
	"net/http"
)
//获取服务器的一些信息
func GetServer(c *gin.Context)  {
	server,err := GlobalService.GetServer()
	if err !=nil{
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.ERROR,
			"msg":   consts.GetMsg(consts.ERROR),
			"data":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  consts.SUCCESS,
		"msg":   consts.GetMsg(consts.SUCCESS),
		"data":  server,
	})
}

type WeatherStruct struct {
	City interface{} `json:"city"`
	List interface{} `json:"list"`
}
type WeatherRequest struct {
	City string `json:"city"`
}
//获取服务器的一些信息
func Weather(c *gin.Context)  {
	var result WeatherStruct
	var weatherRequest WeatherRequest
	//没有错误
	if c.BindJSON(&weatherRequest) == nil {
		url := "http://api.openweathermap.org/data/2.5/forecast/daily?q="+weatherRequest.City+"&mode=json&units=metric&cnt=7&appid=f12159c1f548ea9ab7b5ff1907b1df50"
		weather := utils.Get(url)
		if err := json.Unmarshal([]byte(weather), &result); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  consts.ERROR,
				"msg":   consts.GetMsg(consts.ERROR),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  consts.SUCCESS,
			"msg":   consts.GetMsg(consts.SUCCESS),
			"data":  result,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.INVALID_PARAMS,
			"msg":  consts.GetMsg(consts.INVALID_PARAMS),
		})
	}


}