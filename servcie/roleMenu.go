package servcie

//根据role_id删除对应的menu
func (s *Service) DelUserMenuByRoleId(id int) bool {
	return  s.dao.DelUserMenu(id)
}
