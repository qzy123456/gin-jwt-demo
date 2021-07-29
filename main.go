package main

import (
	"context"
	"github.com/fvbock/endless"
	"flag"
	"fmt"
	"jwtDemo/conf"
	"jwtDemo/routers"
	"jwtDemo/servcie"
	"log"
	"syscall"
)

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
