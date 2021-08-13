package servcie

import "jwtDemo/model"

//获取所有日志
func (s *Service) FindAllLog(pageInfo model.Page) (user []model.Operation) {
	var users []model.Operation
	users, err := s.dao.GetLogByPage(pageInfo)
	if err != nil {
		return nil
	}
	return users
}

//获取所有用户
func (s *Service) GetLogCount(pageInfo model.Page) int64 {
	return s.dao.GetLogCount(pageInfo)
}
