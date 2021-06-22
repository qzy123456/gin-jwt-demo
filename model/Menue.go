package model

//  `menu_id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '权限ID',
//  `menuname` varchar(50) DEFAULT NULL COMMENT '权限名',
//  `menu_url` varchar(50) DEFAULT NULL COMMENT '菜单url',
//  `parent_id` bigint(11) DEFAULT NULL COMMENT '父级ID',
type Menu struct {
	MenuId   int    `json:"menu_id"`
	Menuname string  `json:"menuname"`
	MenuUrl  string  `json:"menu_url"`
	ParentId int    `json:"parent_id"`
}

func (Menu) TableName()string  {
	return "tbl_menu"
}