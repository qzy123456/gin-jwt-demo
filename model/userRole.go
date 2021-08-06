package model

//  `user_id` bigint(11) DEFAULT NULL COMMENT '用户ID',
//  `role_id` bigint(11) NOT NULL COMMENT '角色ID',
type UserRole struct {
	UserId int ` json:"user_id"`
	RoleId int ` json:"role_id"`
	Role  Role 	`xorm:"extends"`	//角色
	RoleMenu RoleMenu `xorm:"extends"`  //菜单
	Menu Menu `xorm:"extends"`
}

type UserRoleNew struct {
	UserId int ` json:"user_id"`
	RoleId int ` json:"role_id"`
}

type Tree struct {
	RoleId int ` json:"role_id"`
	MenuId  int    `json:"menu_id"`
	MenuName string `json:"menu_name"`
	MenuUrl string `json:"menu_url"`
	Children []Tree `json:"children"`
}

func (UserRole) TableName()string  {
	return "tbl_user_role"
}
func (UserRoleNew) TableName()string  {
	return "tbl_user_role"
}
