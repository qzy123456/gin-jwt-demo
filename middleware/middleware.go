package middleware

import (
	"jwtDemo/servcie"
)

type Middleware struct {
	Service *servcie.Service
	NoLoginAction map[string]bool
	NoCheckAction map[string]bool
}

// New middleware service
func New(s *servcie.Service) *Middleware {
	mid := InitNoLoginAction()
	noCheck := InitNoCheckAction()
	return &Middleware{
		Service: s,
		NoLoginAction: mid,
		NoCheckAction: noCheck,
	}
}

//初始化不需要日志的接口
func InitNoLoginAction() map[string]bool  {
	mid := make(map[string]bool)
	mid["/login"] = true
	mid["/server/weather"] = true
	mid["/log/all"] = true
    return mid
}

//初始化不需要检查权限
func InitNoCheckAction() map[string]bool  {
	mid := make(map[string]bool)
	mid["/login"] = true
	mid["/server/weather"] = true
	mid["/user/getMenuById"] = true
	mid["/user/updatePass"] = true
	return mid
}


