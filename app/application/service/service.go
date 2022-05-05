package service

import (
	"vending/app/domain/dto"
	"vending/app/domain/vo"
)

type AuthSrv interface {
	Login(*dto.LoginRe) (*vo.UserVo, error)
	Register(*dto.RegisterRe) (string, error)
}

type CommoditySrv interface {
	// 上架商品
	Up([]string) error
	// 下架商品
	Down([]string) error
	// 导入商品
	Import(string) (count int, err error)
}

type OrderSrv interface {
	// CreateOrder 创建订单 返回付款url
	CreateOrder(interface{}) (string, error)
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
