package dao

import (
	"jwtDemo/model"
)

func (s *Dao)FindAllUser()[]model.User  {
	var users []model.User
	s.Db.Find(&users)
	return users
}


