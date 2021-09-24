package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"jwtDemo/conf"
	_ "jwtDemo/docs"
	"jwtDemo/routers"
	"jwtDemo/servcie"
	"log"
	"syscall"
)

// @title API文档
// @version 1.0.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
func main() {
	flag.Parse()
	// 读取配置
	if err := conf.Init(); err != nil {
		fmt.Println("configs file wrong !!!")
		panic(err)
	}
	// 初始化服务
	svc := servcie.New(&conf.Conf)
	if err := svc.Ping(context.Background()); err != nil {
		panic(err)
	}
	r := routers.InitRouters(svc)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	/*
	启动服务器
	*/
	address := fmt.Sprintf("%s:%d", svc.Conf.Server.Address, svc.Conf.Server.HttpPort)
	server := endless.NewServer(address, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	log.Printf("ListenAndServer At: %s", address)
	// 处理服务器错误
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
