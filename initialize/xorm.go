package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	GroupId int64     `xorm:"index"`
}
type Group struct {
	Id   int64
	Name string
}

//两表连查
type UserGroup struct {
	User  `xorm:"extends"`
	Group `xorm:"extends"`
}

func (UserGroup) TableName() string {
	return "user"
}

type Type struct {
	Id   int64
	Name string
}

//3表联查
type UserGroupType struct {
	User  `xorm:"extends"`
	Group `xorm:"extends"`
	Type  `xorm:"extends"`
}

var engine *xorm.Engine

// Gorm 初始化数据库并产生数据库全局变量
func Xorm() *xorm.Engine {

	host := "127.0.0.1"
	port := "3306"
	username := "root"
	password := "root"
	database := "db_rbac"
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", username, password, host, port, database)
	var err error
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	engine.ShowSQL(true)
    return engine
}


func main() {
	host := "192.168.16.51"
	port := "3306"
	username := "root"
	password := "root"
	database := "test"
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", username, password, host, port, database)
	var err error
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}

	//建表
	err = engine.Sync2(new(User), new(Group), new(Type))
	if err != nil {
		fmt.Println("error in create table user, ", err)
	}

	//新增
	//add()
	//删除
	//dele()
	//修改
	//updat("齐3")
	//查找
	//sele()
	//查一条
	//seleOne()
	//连表
	mergeSelectTwo()
	//3连
	//mergeSelectThree()
}

func add() {
	var user User
	user.Name = "guigui"
	user.Age = 18
	user.Passwd = "23"
	affected, err := engine.Insert(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}

func dele() {
	var user User
	affected, err := engine.ID(2).Delete(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}

func updat(name string) {
	//mm, err := engine.Exec("update user set name = ? where id = ?", name, 1)
	//或者
	var user = new(User)
	user.Name = name
	mm, err := engine.Where("id = ?", 1).Update(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mm)
}

func sele() {

	//results, _ := engine.QueryString("select * from user")
	//for k, v := range results {
	//	fmt.Println(k)
	//	for m, n := range v {
	//		fmt.Println(m)
	//		fmt.Println(n)
	//	}
	//	fmt.Println("=====")
	//}
	engine.ShowSQL(true)
	var users []User
	err := engine.Table("user").Find(&users)
	//以上相当于sql语句：SELECT * FROM user WHERE name = "lgr" AND age > 10
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}

func seleOne() {

	//results, _ := engine.QueryString("select * from user")
	//for k, v := range results {
	//	fmt.Println(k)
	//	for m, n := range v {
	//		fmt.Println(m)
	//		fmt.Println(n)
	//	}
	//	fmt.Println("=====")
	//}
	var users User
	_, err := engine.Where("id = ?", 3).Get(&users)
	//以上相当于sql语句：SELECT * FROM user WHERE name = "lgr" AND age > 10
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}

func mergeSelectTwo() {
	engine.ShowSQL(true)
	users := make([]UserGroup, 0)
	//engine.Join("INNER", "`group`", "group.id = user.group_id").Find(&users)
	engine.SQL("select * from `user`, `group` where user.group_id = group.id").Find(&users)
	fmt.Println(users)
}
func mergeSelectThree() {
	engine.ShowSQL(true)
	users := make([]UserGroupType, 0)
	engine.Table("user").Join("INNER", "`group`", "group.id = user.group_id").
		Join("INNER", "`type`", "type.id = user.group_id").
		Find(&users)
	fmt.Println(users)
	//同时，在使用Join时，也可同时使用Where和Find的第二个参数作为条件，
	// Find的第二个参数同时也允许为各种bean来作为条件。Where里可以是各个表的条件，Find的第二个参数只是被关联表的条件。
	//var name = "齐"
	//engine.Table("user").Join("INNER", "group", "group.id = user.group_id").
	//	Join("INNER", "type", "type.id = user.type_id").
	//	Where("user.name like ?", "%"+name+"%").Find(&users, &User{Name:name})
}
