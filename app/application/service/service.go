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
	// Up 上架商品
	Up([]string) error
	// Down 下架商品
	Down([]string) error
	// Delete 删除商品
	Delete([]string) error
	// Import 导入商品
	Import(interface{}) (count int, err error)
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
