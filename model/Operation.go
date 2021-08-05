package model

//Updated可以让您在记录插入或每次记录更新时自动更新数据库中的标记字段为当前时间，
//需要在xorm标记中使用updated标记，如下所示进行标记，
// 对应的字段可以为time.Time或者自定义的time.Time或者int,int64等int类型。
type Operation struct {
	Id int
	Ip string
	Method string
	Path string
	Body string
	Response string
	CreateTime int64 `xorm:"updated createTime"`
}
