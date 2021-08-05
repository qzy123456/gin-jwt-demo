package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/middleware"
	"jwtDemo/servcie"
	"net/http"
)

var GlobalService  *servcie.Service
func InitRouters(srv *servcie.Service) *gin.Engine {
	GlobalService = srv
	r := gin.Default()
	//加载中间件
	mid := middleware.New(GlobalService)
    //跨域,日志
	r.Use(mid.Cors(),mid.OperationRecord())

	//配置加载静态文件夹，用于显示远程图片
	r.StaticFS("/upload", http.Dir("./upload"))
	r.POST("/login", Login)     //登陆，生成token
	r.POST("/refresh", Refresh) //刷新touken
	r.POST("/file-upload", Upload)//上传文件
	r.POST("/file-uploads", Uploads)//批量上传文件
	user := r.Group("/user").Use(mid.JWTAuth())
	{
		user.GET("/dataByTime", GetDataByTime)  //测试token是否正常
		user.POST("/getMenuById", FindMenuById) //根据用户id获取对应的路由tree
		user.POST("/allUser", FindAllUser)      //所有用户
		user.POST("/saveUser", SaveUser)       //插入一个用户
		user.POST("/deleteById", DeleteById)       //删除一个用户
	}




	return r
}
