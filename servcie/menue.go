package servcie

import (
	"jwtDemo/initialize"
	"jwtDemo/model"
)

func MenuById( id int,menu_id []int)[]*model.Menu  {
	users := make([]*model.Menu,0)
	initialize.Xorm().In("menu_id", menu_id).Where("parent_id = ?",id).Find(&users)
	return users
}
func (s *Service) FindAllMenu()[]model.Tree  {
      return s.dao.GetAllPerm()
}
//检测menu重复
func (s *Service)CheckMenuByName(menu model.Menu) bool  {
	return s.dao.CheckMenuByName(menu)
}
//插入一条menu
func (s *Service)SaveMenu(menu model.Menu) error  {
	return s.dao.SaveMenu(menu)
}
//删除一个menu
func (s *Service)DeleteMenuById(id int)error  {
	return s.dao.DeleteMenuById(id)
}
//修改
func (s *Service)UpdateMenuById(menu model.Menu)error  {
	return s.dao.UpdateMenuById(menu)
}
//通过菜单id，拼接成递归样式
func (s *Service) GetAllParentIds(id int)interface{} {
return 	s.dao.GetAllParentIds(id)

}