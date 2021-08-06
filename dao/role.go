package dao

import (
	"jwtDemo/model"
)

//根据用户id 先删除之前的角色
func (s *Dao) DeleteUserRoleById(id int) bool{

	user := new(model.UserRoleNew)
	user.UserId = id
	affected, err := s.Db.Delete(user)
	if err != nil  || affected <= 0{
		return false
	}
	return true
}
