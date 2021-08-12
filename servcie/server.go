package servcie

import (
	"jwtDemo/utils"
)

func (ser *Service) GetServer() (server *utils.Server, err error)  {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil{
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil{
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil{
		return &s, err
	}

	return &s, nil
}
