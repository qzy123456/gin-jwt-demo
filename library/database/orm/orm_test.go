package orm

import (
	"fmt"
	"jwtDemo/model"
	"testing"
)

func TestNewMySQL(t *testing.T) {
	c := &Config{
		Dsn: []string{"root:root@tcp(127.0.0.1:3306)/db_rbac?charset=utf8","root:root@tcp(127.0.0.1:3306)/db_rbac?charset=utf8"},
	}

	db := NewMySQL(c)

	var users []model.User
	db.Find(&users)
	fmt.Println(users)
}
