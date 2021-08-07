package servcie

import "jwtDemo/model"

//根据role_id删除对应的menu
func (s *Service) DelUserMenuByRoleId(id int) bool {
	return  s.dao.DelUserMenu(id)
}

func (s *Service) SaveRoleMenu(menu model.RoleMenuNew)error  {
	return  s.dao.SaveRoleMenu(menu)
}