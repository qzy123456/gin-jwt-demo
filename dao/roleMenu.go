package dao

import (
	"fmt"
	"jwtDemo/model"
)

//根据id 删除menu
func (s *Dao) DelUserMenu(id int) bool{

	user := new(model.RoleMenu)
	user.RoleId = id
	affected, err := s.Db.Delete(user)
	fmt.Println(err)
	if err != nil  || affected <= 0{
		return false
	}
	return true
}