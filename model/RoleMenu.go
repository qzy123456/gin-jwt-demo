package model


// `role_id` bigint(11) NOT NULL COMMENT '角色ID',
//  `menu_id` bigint(11) DEFAULT NULL COMMENT '权限ID',
type RoleMenu struct {
    RoleId int  ` json:"role_id"`
    MenuId int  ` json:"menu_id"`
}

type RoleMenuNew struct {
    RoleId int     `json:"role_id"`
    MenuId string  `json:"menu_id"`
}

func (RoleMenu) TableName()string  {
    return "tbl_role_menu"
}

func (RoleMenuNew) TableName()string  {
    return "tbl_role_menu"
}