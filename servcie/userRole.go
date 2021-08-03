package servcie

import (
	"fmt"
	"jwtDemo/initialize"
	"jwtDemo/model"
)

//查询所有的当前用户权限的菜单ID
func FindMenuById( id int)[]model.UserRole  {
	users := make([]model.UserRole, 0)
	initialize.Xorm().Select("*").Table("tbl_user_role").Join("INNER", "`tbl_role`", "tbl_user_role.user_id= tbl_role.role_id").
		Join("INNER", "`tbl_role_menu`", "tbl_user_role.role_id = tbl_role_menu.role_id").
		Join("INNER", "`tbl_menu`", "tbl_role_menu.menu_id = tbl_menu.menu_id").
		Where("tbl_user_role.role_id = ?",id).
		Find(&users)
	return users
}

//查询所有的当前用户权限的菜单ID
func FindMenus( id int)[]int  {
	users := make([]int, 0)
	initialize.Xorm().Select("tbl_role_menu.menu_id").Table("tbl_user_role").Join("INNER", "`tbl_role`", "tbl_user_role.user_id= tbl_role.role_id").
		Join("INNER", "`tbl_role_menu`", "tbl_user_role.role_id = tbl_role_menu.role_id").
		//Join("INNER", "`tbl_menu`", "tbl_role_menu.menu_id = tbl_menu.menu_id").
		Where("tbl_user_role.role_id = ?",id).
		Find(&users)
	return users
}
//通过菜单id，拼接成递归样式
func (s *Service) GetAllPerm2(id int)([]model.Tree) {
	data := FindMenuById(id)
	var options []model.Tree
	o := GetAllPerm3(data,0,options)
	return o
}
//通过菜单id，拼接成递归样式(这种最优，只需要查询一次数据库)
func GetAllPerm3(data []model.UserRole,id int,cp []model.Tree) []model.Tree {
	var options []model.Tree

	for _, value := range data {
		if value.Menu.ParentId == id{
			var option model.Tree
			option.RoleId = value.RoleId
			option.MenuId = value.Menu.MenuId
			option.MenuName = value.Menu.Menuname
			option.MenuUrl = value.Menu.MenuUrl
			option.Children = GetAllPerm3(data,value.Menu.MenuId,options)
			cp = append(cp,option)
		}
	}
	return  cp
}
//通过菜单id，拼接成递归样式（这种可以，非递归，查数据）
func GetAllPerm(ids []int)([]model.Tree) {
	var menu []*model.Menu
	//得到所有的0，也就是根节点
	menu = MenuById(0,ids)
	fmt.Println("menu",menu)
	treeList := []model.Tree{}
	for _, v := range menu { // 循环所有父菜单
		childs := MenuById(v.MenuId,ids) // 拿到每个父菜单的子菜单
		childss := make([]model.Tree,0)
		for _,vv := range childs{
			child := model.Tree{ // 拼装父菜单下面的子菜单
				RoleId:vv.MenuId,
				MenuId:vv.MenuId,
				MenuName:vv.Menuname,
				MenuUrl:vv.MenuUrl,
			}
			childss = append(childss,child)
		}
		node := model.Tree{ // 拼装父菜单数据
			RoleId:v.MenuId,
			MenuId:v.MenuId,
			MenuName:v.Menuname,
			MenuUrl:v.MenuUrl,
		}
		node.Children = childss
		treeList = append(treeList, node)
	}
	return treeList
}
//通过菜单id，拼接成递归样式(这种可以，递归，查数据库)
func GetAllPerm4(id int,ids []int)([]model.Tree) {
	var menu []*model.Menu
	//得到所有的0，也就是根节点
	menu = MenuById(id,ids)
	fmt.Println("menu",menu)
	treeList := []model.Tree{}
	for _, v := range menu { // 循环所有父菜单
		node := model.Tree{ // 拼装父菜单数据
			RoleId:v.MenuId,
			MenuId:v.MenuId,
			MenuName:v.Menuname,
			MenuUrl:v.MenuUrl,
		}
		node.Children = GetAllPerm4(v.MenuId,ids)
		treeList = append(treeList, node)
	}
	return treeList
}
