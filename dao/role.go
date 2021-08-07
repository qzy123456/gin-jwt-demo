package dao

import (
	"fmt"
	"jwtDemo/model"
)

//根据用户id 先删除之前的角色
func (s *Dao) DeleteUserRoleById(id int) bool{

	user := new(model.UserRoleNew)
	user.UserId = id
	affected, err := s.Db.Delete(user)
	if err != nil  || affected <= 0{
		return false
	}
	return true
}

//查询用户总量，用于分页
func (s *Dao) GetRoleCount(pageInfo model.Page) int64  {
	user := new(model.Role)
	total, _ := s.Db.Where("role_name like ? ","%"+pageInfo.Query+"%").Count(user)
	return  total
}

//分页查询用户总量
func (s *Dao) GetRoleByPage(pageInfo model.Page) (users []model.Role,errs error)  {
	var user  []model.Role
	err := s.Db.
		Where("role_name like ? ","%"+pageInfo.Query+"%").Limit(pageInfo.PageSize, (pageInfo.PageNum - 1) * pageInfo.PageSize).Find(&user)
	fmt.Println(err)
	s.Db.ShowSQL(true)
	if err!= nil{
		return nil,err
	}
	return user,err
}
//检测用户名是否存在
func (s *Dao) CheckRoleByUserName(username string) bool {
	role := new(model.Role)
	has, err := s.Db.Where("username=?", username).Count(role)
	if err!= nil{
		return false
	}
	if has >= 1{
		return true
	}
	return false
}
//新增一个角色
func (s *Dao) SaveRole(user model.Role) error{
	_, err := s.Db.Insert(&user)
	if err != nil {
		return err
	}
	return  nil
}
//根据id删除角色
func (s *Dao) DeleteRoleById(id int) bool {
	user := new(model.Role)
	user.RoleId = id
	affected, err := s.Db.Delete(user)
   fmt.Println(err)
	if err != nil  || affected <= 0{
		return false
	}
	return true
}
//更新角色信息
func (s *Dao) UpdateRoleById(us model.Role) bool {
	user := new(model.Role)
	user = &us
	affected, err := s.Db.Where("role_id = ?", user.RoleId).Update(user)
	if err != nil  || affected < 0{
		return false
	}
	return true
}