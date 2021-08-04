package dao

import (
	"fmt"
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

//查询用户总量，用于分页
func (s *Dao) GetUserCount(pageInfo model.Page) int64  {
	user := new(model.User)
	total, _ := s.Db.Where("username like ? ","%"+pageInfo.Query+"%").Count(user)
	return  total
}

//分页查询用户总量
func (s *Dao) GetUserByPage(pageInfo model.Page) (users []model.User,errs error)  {
	var user  []model.User
	fmt.Println(pageInfo)
	err := s.Db.Where("username like ? ","%"+pageInfo.Query+"%").Limit(pageInfo.PageSize, (pageInfo.PageNum - 1) * pageInfo.PageSize).Find(&user)
	if err!= nil{
		return nil,err
	}
	return user,err
}


