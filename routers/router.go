package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/api"
	"jwtDemo/middleware/jwt"
	"net/http"
)

func InitRouters() *gin.Engine {

	r := gin.Default()

	//配置加载静态文件夹，用于显示远程图片
	r.StaticFS("/upload", http.Dir("./upload"))
	r.POST("/login", api.Login)
	r.POST("/register", api.RegisterUser)
	r.POST("/refresh", api.Refresh) //刷新touken
	//上传文件
	r.POST("/file-upload", api.Upload)
	r.POST("/file-uploads", api.Uploads)
	taR := r.Group("/data")
	taR.Use(jwt.JWTAuth())
	{
		taR.GET("/dataByTime", api.GetDataByTime)
	}

	r.POST("/allUser", api.FindAllUser)
	r.POST("/getMenuById", api.FindMenuById)

	return r
}
