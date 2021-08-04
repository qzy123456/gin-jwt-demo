package middleware

import (
	"jwtDemo/servcie"
)

type Middleware struct {
	Service *servcie.Service
	NoLoginAction map[string]bool
}

// New middleware service
func New(s *servcie.Service) *Middleware {
	mid := InitNoLoginAction()
	return &Middleware{
		Service: s,
		NoLoginAction: mid,
	}
}

//初始化不需要日志的接口
func InitNoLoginAction() map[string]bool  {
	mid := make(map[string]bool)
	mid["/login"] = true
	mid["/file-upload"] = true
    return mid
}
