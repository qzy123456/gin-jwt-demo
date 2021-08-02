package dao

import (
	"jwtDemo/model"
)

func (s *Dao) SaveOperation(operation model.Operation) error{
	_, err := s.Db.Insert(&operation)
	if err != nil {
		return err
	}
	return  nil
}
