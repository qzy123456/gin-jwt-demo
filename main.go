package main

import (
	"context"
	"flag"
	"fmt"
	"jwtDemo/conf"
	"jwtDemo/routers"
	"jwtDemo/servcie"
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
	_ = r.Run(":8080")
}
