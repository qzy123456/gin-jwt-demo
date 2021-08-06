package servcie

import (
	"jwtDemo/model"
	"jwtDemo/utils"
)

//获取所有用户
func (s *Service) FindAllUser(pageInfo model.Page) (user []model.User) {
	var users []model.User
	users, err := s.dao.GetUserByPage(pageInfo)
	if err != nil {
		return nil
	}
	return users
}

//获取所有用户
func (s *Service) GetUserCount(pageInfo model.Page) int64 {
	return s.dao.GetUserCount(pageInfo)
}

//检测用户的账号，密码
func (s *Service) CheckLogin(loginReq model.LoginReq) *model.User {
	user := new(model.User)
	useInfo := s.dao.CheckLogin(loginReq)
	if useInfo != nil {
		user = useInfo
		return user
	}
	return nil
}

//检测用户的账号，密码
func (s *Service) SaveUser(user model.User) error {
	user.CreateTime = utils.GetYmd()
	user.LastTime = utils.GetYmds()
	err := s.dao.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

//检测用户的账号，密码
func (s *Service) CheckUserByName(username string) bool {
	useInfo := s.dao.CheckUserByUserName(username)
	if useInfo {
		return true
	}
	return false
}

//根据id删除
func (s *Service) DeleteById(id int) bool {
	useInfo := s.dao.DeleteById(id)
	if useInfo {
		return true
	}
	return false
}

//根据id修改
func (s *Service) UpdateById(user model.User) bool {
	useInfo := s.dao.UpdateById(user)
	if useInfo {
		return true
	}
	return false
}

//给用户分配角色
func (s *Service) SaveUserRole(user model.UserRoleNew) error {
	err := s.dao.SaveUserRole(user)
	if err != nil {
		return err
	}
	return nil
}