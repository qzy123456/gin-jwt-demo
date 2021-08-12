package dao

import (
	"fmt"
	"jwtDemo/model"
)


//menu
func (s *Dao) FindAllMenu() (user []model.UserRole)  {
	users := make([]model.UserRole, 0)
	err :=  s.Db.Select("*").Table("tbl_user_role").Join("INNER", "`tbl_role`", "tbl_user_role.user_id= tbl_role.role_id").
		Join("INNER", "`tbl_role_menu`", "tbl_user_role.role_id = tbl_role_menu.role_id").
		Join("INNER", "`tbl_menu`", "tbl_role_menu.menu_id = tbl_menu.menu_id").
		Find(&users)
	fmt.Println(err)
	return users
}
//通过菜单id，拼接成递归样式
func (s *Dao)GetAllPerm()([]model.Tree) {
	data := s.FindAllMenu()
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
			option.IsShow =  value.Menu.IsShow
			option.Children = GetAllPerm3(data,value.Menu.MenuId,options)
			cp = append(cp,option)
		}
	}
	return  cp
}
//查询menu是否重复
func (s *Dao) CheckMenuByName(menus model.Menu) bool {
	menu := new(model.Menu)
	var has int64
	var err error
	if menus.MenuId != 0 {
		has, err = s.Db.Where("menu_url=?", menus.MenuUrl).NotIn("menu_id", menus.MenuId).Count(menu)
	} else {
		has, err = s.Db.Where("menu_url=?", menus.MenuUrl).Count(menu)

	}
	if err != nil {
		return false
	}
	if has >= 1 {
		return true
	}

	return false
}
//插入一条menu
func (s *Dao)SaveMenu(menu model.Menu) error  {
	_, err := s.Db.Insert(&menu)
	if err != nil {
		return err
	}
	return  nil
}
//删除
func (s *Dao)DeleteMenuById(id int) error{
	_, err := s.Db.Exec("DELETE FROM `tbl_menu` WHERE (((menu_id = ?)) OR (parent_id = ?))", id,id)
	if err != nil  {
		return err
	}
	return nil
}
//修改
//因为go的结构初始值都是各自的0值，所以xorm无法识别要更新结构体中的哪个col，Cols方法写可以正常解决。
func (s *Dao)UpdateMenuById(menu model.Menu)error  {
	user := new(model.Menu)
	user = &menu
	_, err := s.Db.Where("menu_id = ?", user.MenuId).
		Cols("menuname,menu_url,parent_id,is_show").Update(user)
	if err != nil {
		return err
	}
	return nil
}
//查询所有的路由
func (s *Dao) FindMenuByRoleIds()[]model.MenuNew{
	users := make([]model.MenuNew, 0)
	s.Db.Select("*").Table("tbl_menu").
		Find(&users)
	return users
}
//根据menuid查询menu信息
func (s *Dao)FindMenuByMenuIds(id int) (menu model.Menu)  {
	s.Db.Where("menu_id = ?",id).Get(&menu)
	return menu
}

func (s *Dao)GetAllParentIds(id int) interface{}  {
	var sql = "select  T1.menu_id from tbl_menu T1 inner join ( select @id as _id, (select @id := parent_id from tbl_menu where menu_id = _id) AS pid from (select @id := ?) T2_1,tbl_menu T2_2 where @id is not null and @id != '') T2 on T1.parent_id = T2._id"
	res , err := s.Db.QueryString(sql,id)
	fmt.Println(res,err)
	return res
}
