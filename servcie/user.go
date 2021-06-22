package servcie

import (
	"jwtDemo/initialize"
	"jwtDemo/model"
)

func FindAllUser()[]model.User  {
	var users []model.User
	initialize.Gorm().Find(&users)
	return users
}

