package dao

import (
	"fmt"
	"jwtDemo/model"
	"jwtDemo/utils"
)
//检测登陆的账号密码
func (s *Dao) CheckLogin(loginReq model.LoginReq) *model.User  {
	user := new(model.User)
	has, err := s.Db.Where("username=?", loginReq.Username).Where("password=?", loginReq.Password).Get(user)
	if !has || err!= nil{
		return nil
	}
	return user
}

//查询用户总量，用于分页
func (s *Dao) GetUserCount(pageInfo model.Page) int64  {
	user := new(model.User)
	total, _ := s.Db.Where("username like ? ","%"+pageInfo.Query+"%").Count(user)
	return  total
}

//分页查询用户总量
func (s *Dao) GetUserByPage(pageInfo model.Page) (users []model.User,errs error)  {
	var user  []model.User
	err := s.Db.Table(&model.User{}).Join("LEFT", "tbl_user_role", "tbl_user_role.user_id = tbl_user.user_id").
		Join("LEFT", "tbl_role", "tbl_user_role.role_id = tbl_role.role_id").
		Where("tbl_user.username like ? ","%"+pageInfo.Query+"%").Limit(pageInfo.PageSize, (pageInfo.PageNum - 1) * pageInfo.PageSize).Find(&user)
	fmt.Println(err)
	s.Db.ShowSQL(true)
	if err!= nil{
		return nil,err
	}
	return user,err
}
//新增一个用户
func (s *Dao) SaveUser(user model.UserNew) error{
	_, err := s.Db.Insert(&user)
	if err != nil {
		return err
	}
	return  nil
}
//检测用户名是否存在
func (s *Dao) CheckUserByUserName(users model.UserNew) bool {
	user := new(model.UserNew)
	var has int64
	var err error
	if users.UserId != 0 {
		has, err = s.Db.Where("username=?", users.Username).NotIn("userId", users.UserId).Count(user)
	} else {
		has, err = s.Db.Where("username=?", users.Username).Count(user)

	}
	if err != nil {
		return false
	}
	if has >= 1 {
		return true
	}

	return false
}
//检测登陆的账号密码
func (s *Dao) DeleteById(id int) bool {
	user := new(model.User)
	user.UserId = id
	affected, err := s.Db.Delete(user)
    fmt.Println(user)
	if err != nil  || affected <= 0{
	 return false
	}
	return true
}
//检测登陆的账号密码
func (s *Dao) UpdateById(us model.UserNew) bool {
	user := new(model.UserNew)
	user = &us
	user.LastTime = utils.GetYmds()
	affected, err := s.Db.Where("user_id = ?", user.UserId).
		Cols("username,password,enabled,last_time").Update(user)
	if err != nil  || affected < 0{
		return false
	}
	return true
}
//新增一个用户
func (s *Dao) SaveUserRole(user model.UserRoleNew) error{
	_, err := s.Db.Insert(user)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return  nil
}
//删除用户顺便删除权限
func (s *Dao) DelUserRole(id int) error{
	user := new(model.UserRoleNew)
	user.UserId = id
	_, err := s.Db.Delete(user)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return  nil
}
