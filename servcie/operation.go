package servcie

import (
	"jwtDemo/model"
)

func (s *Service) SaveOperation(operation model.Operation) {
	 s.dao.SaveOperation(operation)
}
