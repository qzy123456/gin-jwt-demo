package servcie

//根据用户id 先删除之前的角色
func (s *Service) DeleteUserRoleById(id int) bool{
	return s.dao.DeleteUserRoleById(id)
}
