package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/middleware/jwt"
	"jwtDemo/servcie"
	"net/http"
)

var GlobalService  *servcie.Service
func InitRouters(srv *servcie.Service) *gin.Engine {
	GlobalService = srv
	r := gin.Default()

	//配置加载静态文件夹，用于显示远程图片
	r.StaticFS("/upload", http.Dir("./upload"))
	r.POST("/login", Login)
	r.POST("/register", RegisterUser)
	r.POST("/refresh", Refresh) //刷新touken
	//上传文件
	r.POST("/file-upload", Upload)
	r.POST("/file-uploads", Uploads)
	taR := r.Group("/data")
	taR.Use(jwt.JWTAuth())
	{
		taR.GET("/dataByTime", GetDataByTime)
	}

	r.POST("/allUser", FindAllUser)
	r.POST("/getMenuById", FindMenuById)

	return r
}
