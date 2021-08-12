package model


type Menu struct {
	MenuId   int    `json:"menu_id" xorm:"pk menu_id autoincr"`
	Menuname string  `json:"menu_name" xorm:"menuname"`
	MenuUrl  string  `json:"menu_url" xorm:"menu_url"`
	ParentId int    `json:"parent_id" xorm:"parent_id"`
	IsShow   int    `json:"isShow" xorm:"is_show"`
}
type MenuNew struct {
	MenuId   int    `json:"menu_id" xorm:"pk menu_id autoincr"`
	Menuname string  `json:"menu_name" xorm:"menuname"`
	MenuUrl  string  `json:"menu_url" xorm:"menu_url"`
	ParentId int    `json:"parent_id" xorm:"parent_id"`
	IsShow   int    `json:"isShow" xorm:"is_show"`
	Children []MenuNew `json:"children" xorm:"-"`
}
type Menus struct {
	MenuId   int    `json:"menu_id" xorm:"pk menu_id autoincr"`
	Menuname string  `json:"menu_name" xorm:"menuname"`
	MenuUrl  string  `json:"menu_url" xorm:"menu_url"`
	ParentId int    `json:"parent_id" xorm:"parent_id"`
	IsShow   int    `json:"isShow" xorm:"is_show"`
	ParentIds []int `json:"parent_ids" xorm:"-"`
}

func (Menu) TableName()string  {
	return "tbl_menu"
}
func (Menus) TableName()string  {
	return "tbl_menu"
}