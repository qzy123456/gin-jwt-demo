package servcie

import (
"fmt"
"jwtDemo/model"
)

func (s *Service) FindAllUser() []model.User {
	var users []model.User
	err := s.dao.Db.Find(&users)
	if err != nil {
		fmt.Println(err)
	}
	return users
}
