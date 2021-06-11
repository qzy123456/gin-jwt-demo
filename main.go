package main

import (
	"github.com/gin-gonic/gin"

	"jwtDemo/api"
	"jwtDemo/middleware/jwt"
)

func main() {
	r := gin.Default()
	r.POST("/login", api.Login)
	r.POST("/register", api.RegisterUser)
	r.POST("/refresh", api.Refresh) //刷新touken

	taR := r.Group("/data")
	taR.Use(jwt.JWTAuth())

	{
		taR.GET("/dataByTime", api.GetDataByTime)
	}
	r.Run(":8080")
}
