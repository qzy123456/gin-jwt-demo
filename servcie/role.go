package servcie

import (
	"jwtDemo/model"
)

//根据用户id 先删除之前的角色
func (s *Service) DeleteUserRoleById(id int) bool{
	return s.dao.DeleteUserRoleById(id)
}
//获取所有用户
func (s *Service) FindAllRole(pageInfo model.Page) (user []model.Role) {
	var users []model.Role
	users, err := s.dao.GetRoleByPage(pageInfo)
	if err != nil {
		return nil
	}
	return users
}
//获取所有用户
func (s *Service) GetRoleCount(pageInfo model.Page) int64 {
	return s.dao.GetRoleCount(pageInfo)
}
//检测用户的账号，密码
func (s *Service) CheckRoleByName(username string) bool {
	useInfo := s.dao.CheckRoleByUserName(username)
	if useInfo {
		return true
	}
	return false
}
//插入一条数据
func (s *Service) SaveRole(user model.Role) error {
	err := s.dao.SaveRole(user)
	if err != nil {
		return err
	}
	return nil
}
//根据id删除
func (s *Service) DeleteRoleById(id int) bool {
	useInfo := s.dao.DeleteRoleById(id)
	if useInfo {
		return true
	}
	return false
}

//根据id修改
func (s *Service) UpdateRoleById(user model.Role) bool {
	useInfo := s.dao.UpdateRoleById(user)
	if useInfo {
		return true
	}
	return false
}