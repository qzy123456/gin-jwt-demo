package dao

import (
	"fmt"
	"jwtDemo/model"
	"strconv"
	"strings"
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
//插入角色对应的menu
func (s *Dao) SaveRoleMenu(menu model.RoleMenuNew) error{
	var menus model.RoleMenu
	var err  error
	var MenuId  int
	menuIds := strings.Split(menu.MenuId,",")
	for _, value := range menuIds {
		MenuId,err = strconv.Atoi(value)
		menus.MenuId = MenuId
		menus.RoleId = menu.RoleId
		_, err = s.Db.Insert(menus)
		if err != nil {
			return err
		}
	}
    return nil

}