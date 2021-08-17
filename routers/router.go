package routers

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/middleware"
	"jwtDemo/servcie"
	"net/http"
)

var GlobalService *servcie.Service

func InitRouters(srv *servcie.Service) *gin.Engine {
	GlobalService = srv
	r := gin.Default()
	//加载中间件
	mid := middleware.New(GlobalService)
	//跨域,日志,错误，检测权限
	r.Use(mid.Cors(), mid.SetUp())

	//配置加载静态文件夹，用于显示远程图片
	r.StaticFS("/upload", http.Dir("./upload"))
	r.POST("/login", Login)          //登陆，生成token
	r.POST("/refresh", Refresh)      //刷新touken
	r.POST("/file-upload", Upload)   //上传文件
	r.POST("/file-uploads", Uploads) //批量上传文件
	//用户
	user := r.Group("/user").Use(mid.JWTAuth(),mid.CheckMenus(),mid.OperationRecord(),)
	{
		user.GET("/dataByTime", GetDataByTime)  //测试token是否正常
		user.POST("/getMenuById", FindMenuById) //根据用户id获取对应的路由tree
		user.POST("/allUser", FindAllUser)      //所有用户
		user.POST("/saveUser", SaveUser)        //插入一个用户
		user.POST("/deleteById", DeleteById)    //删除一个用户
		user.POST("/update", UpdateById)        //修改一个用户
		user.POST("/insertRole", InsertRole)    //给用户分配角色
		user.POST("/updateStatus", UpdateStatus)    //给用户分配角色
		user.POST("/updatePass", UpdatePass)     //更改密码
	}
	//角色
	role := r.Group("/role").Use(mid.JWTAuth(),mid.CheckMenus(),mid.OperationRecord(),)
	{
		role.POST("/all", FindAllRole)   //所有用户
		role.POST("/save", SaveRole)     //插入一个角色
		role.POST("/delete", DeleteRoleById) //删除一个角色
		role.POST("/update", UpdateroleById)     //修改一个角色
		role.POST("/getMenuById", GetMenuById)     //根据roleid，返回所有的menu信息
		role.POST("/deleteMenuAndRoleId", DeleteMenuAndRoleId)     //根据roleid，返回所有的menu信息
	}
	//菜单
	menu := r.Group("/menu").Use(mid.JWTAuth(),mid.CheckMenus(),mid.OperationRecord(),)
	{
		menu.POST("/all", FindAllMenu)   //所有菜单
		menu.POST("/save", SaveMenu)     //插入一个菜单
		menu.POST("/delete", DeleteMenuById) //删除一个菜单
		menu.POST("/update", UpdateMenuById)     //修改一个菜单
		menu.POST("/giveMenu", GiveMenu)     //给角色分配菜单
		menu.POST("/getMenuById", GetMenuByMenuId)  //根据菜单id 查询
	}
	//系统信息
	server := r.Group("/server").Use(mid.JWTAuth(),mid.CheckMenus(),mid.OperationRecord(),)
	{
		server.POST("/server", GetServer)   //所有菜单
		server.POST("/weather", Weather)   //天气
	}
	//日志信息
	log := r.Group("/log").Use(mid.JWTAuth(),mid.CheckMenus(),mid.OperationRecord(),)
	{
		log.POST("/all", GetLogs)   //分页获取所有的日志
	}
	return r
}
