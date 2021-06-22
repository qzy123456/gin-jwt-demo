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