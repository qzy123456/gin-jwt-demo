package model


// `role_id` bigint(11) NOT NULL COMMENT '角色ID',
//  `menu_id` bigint(11) DEFAULT NULL COMMENT '权限ID',
//  KEY `FK_FK_PERMISSION` (`menu_id`),
//  KEY `FK_FK_ROLE` (`role_id`),
type RoleMenu struct {
    RoleId int  ` json:"role_id"`
    MenuId int  ` json:"menu_id"`
    //Menus  Menu  `xorm:"extends"` //菜单
}
