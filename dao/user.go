package dao

import (
	"fmt"
	"github.com/pkg/errors"
	"jwtDemo/model"
	"jwtDemo/utils"
)
//检测登陆的账号密码
func (s *Dao) CheckLogin(loginReq model.LoginReq) *model.UserNew  {
	user := new(model.UserNew)
	has, err := s.Db.Where("username=?", loginReq.Username).Where("password=?", loginReq.Password).Where("enabled=?", 1).Get(user)
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
    fmt.Println(err)
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
	affected, err := s.Db.Where("user_id = ?", user.UserId).Update(user)
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
//检测登陆的账号密码
func (s *Dao) UpdateStatus(us model.UserNew) bool {
	user := new(model.UserNew)
	user = &us
	user.LastTime = utils.GetYmds()
	affected, err := s.Db.Where("user_id = ?", user.UserId).Cols("enabled").Update(user)
	if err != nil  || affected < 0{
		return false
	}
	return true
}

//修改密码
func (s *Dao)UpdatePass(user model.UpdatePass) error  {
	//先查询，用户id对应的密码是否匹配
	use := new(model.UserNew)
	has, err := s.Db.Where("username=?", user.UserName).Get(use)
	if !has || err!= nil{
		return errors.New("用户查不到")
	}
    if use.Password != user.Password{
		return errors.New("密码不匹配")
	}
	if user.NewPassword == user.Password{
		return errors.New("新密码不能等于旧密码")
	}
	uses := new(model.UserNew)
	uses.Username = user.UserName
	uses.Password = user.NewPassword
	//更改密码
	_,err1 := s.Db.Where("username = ?", user.UserName).Cols("password").Update(uses)
	if err1 != nil  {
		return err1
	}
	return nil
}
