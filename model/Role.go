package model

//`role_id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
//  `role_name` varchar(50) DEFAULT NULL COMMENT '角色名称',
type Role struct {
    RoleId int  `json:"role_id" xorm:"pk role_id autoincr"`
    RoleName string  `json:"role_name" xorm:"role_name"`
}
func (Role) TableName()string  {
    return "tbl_role"
}