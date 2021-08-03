package dao

import (
	"jwtDemo/model"
)
//检测登陆的账号密码
func (s *Dao) CheckLogin(loginReq model.LoginReq) *model.User  {
	user := new(model.User)
	has, err := s.Db.Where("username=?", loginReq.Username).Where("password=?", loginReq.Password).Get(user)
	if !has || err!= nil{
		return nil
	}
	return user
}


