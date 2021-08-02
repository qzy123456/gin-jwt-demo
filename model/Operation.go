package model


type Operation struct {
	Id int
	Ip string
	Method string
	Path string
	Body string
	Response string
	CreateTime int64 `xorm:"createTime"`
}
