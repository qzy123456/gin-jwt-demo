package servcie

import (
	"jwtDemo/model"
)
//获取所有用户
func (s *Service) FindAllUser(pageInfo model.Page) (user []model.User)   {
	var users []model.User
	users, err := s.dao.GetUserByPage(pageInfo)
	if err != nil {
		return nil
	}
	return users
}
//获取所有用户
func (s *Service) GetUserCount(pageInfo model.Page) int64  {
	return s.dao.GetUserCount(pageInfo)
}
//检测用户的账号，密码
func (s *Service) CheckLogin(loginReq model.LoginReq) *model.User {
	user := new(model.User)
	useInfo := s.dao.CheckLogin(loginReq)
	if useInfo !=nil{
		user = useInfo
		return user
	}
	return nil
}
