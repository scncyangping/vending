package service

import (
	"vending/app/domain/dto"
	"vending/app/domain/vo"
)

type AuthSrv interface {
	Login(*dto.LoginRe) (*vo.UserVo, error)
	Register(*dto.RegisterRe) (string, error)
}

//
//type SrvManager struct {
//	AuthSrv
//}
//
//func NewAuthSrvManager(srv service.Service) *SrvManager {
//	return &SrvManager{
//		AuthSrv: srv.UserSrv,
//	}
//}
