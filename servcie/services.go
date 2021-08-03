package servcie

import (
	"context"
	"github.com/sirupsen/logrus"
	"jwtDemo/conf"
	"jwtDemo/dao"
	"jwtDemo/library/xlog"
)

type Service struct {
	Conf          *conf.Config
	dao           *dao.Dao
	log           *logrus.Logger
	RequestLogger *logrus.Logger
	HttpLogger    *logrus.Logger
	Jwt           string
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		Conf:          c,
		dao:           dao.New(c),
		log:           xlog.Init(c.Log, "business"),
		RequestLogger: xlog.Init(c.Log, "request"),
		HttpLogger:    xlog.Init(c.Log, "http"),
	}
	return
}

func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}
