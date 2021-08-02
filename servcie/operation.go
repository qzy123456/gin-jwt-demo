package servcie

import (
	"jwtDemo/model"
)

func (s *Service) SaveOperation(operation model.Operation) error {

	err := s.dao.SaveOperation(operation)
	if err != nil {
		return err
	}
	return nil
}
