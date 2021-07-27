package main

import (
	"jwtDemo/routers"
)

func main() {
	r := routers.InitRouters()
	_ = r.Run(":8080")
}
