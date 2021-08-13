package dao

import (
	"jwtDemo/model"
)
//插入日志
func (s *Dao) SaveOperation(operation model.Operation){
	 s.Db.Insert(&operation)
}

//查询用户总量，用于分页
func (s *Dao) GetLogCount(pageInfo model.Page) int64  {
	operation := new(model.Operation)
	total, _ := s.Db.Where("path like ? ","%"+pageInfo.Query+"%").Count(operation)
	return  total
}

//分页查询日志
func (s *Dao) GetLogByPage(pageInfo model.Page) (users []model.Operation,errs error)  {
	var operation  []model.Operation
	err := s.Db.Desc("createTime").Where("path like ? ","%"+pageInfo.Query+"%").
		Limit(pageInfo.PageSize, (pageInfo.PageNum - 1) * pageInfo.PageSize).Find(&operation)
	if err!= nil{
		return nil,err
	}
	return operation,err
}
