package servcie

import (
"fmt"
"jwtDemo/model"
)
//获取所有用户
func (s *Service) FindAllUser() []model.User {
	var users []model.User
	err := s.dao.Db.Find(&users)
	if err != nil {
		fmt.Println(err)
	}
	return users
}
//检测用户是否登陆
func (s *Service) CheckLogin(loginReq model.LoginReq) *model.User {
	user := new(model.User)
	useInfo := s.dao.CheckLogin(loginReq)
	if useInfo !=nil{
		user = useInfo
		return user
	}
	return nil
}
