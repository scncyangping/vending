package service

import (
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/dto"
	"vending/app/application/service/impl"
	"vending/app/domain/service"
)

type AuthSrv interface {
	Login(*cmd.LoginCmd) (dto.UserDto, error)
	Register(*cmd.RegisterCmd) (string, error)
}

//CommoditySrv 商品接口
type CommoditySrv interface {
	// Save 新增商品
	Save(*cmd.CommoditySaveCmd) (string, error)
	// Update 更新商品
	Update(*cmd.CommodityUpdateCmd) error
	// Delete 删除商品
	Delete([]string) error
	// Get 查询商品明细
	Get(string) *dto.CommodityDto
	// Up 商品上架
	Up(string) error
	// Down 商品下架
	Down(string) error
	// Query 商品列表
	Query(query.CommoditiesPageQuery) ([]*dto.CommodityDto, error)
}

// InventorySrv 库存
type InventorySrv interface {
	// SaveCategory 保存库存分类
	SaveCategory(*cmd.CategorySaveCmd) (string, error)
	// UpdateCategory 更新分类
	UpdateCategory(*cmd.CategoryUpdateCmd) error
	// RemoveCategoryByIds 移除分类
	RemoveCategoryByIds([]string) error
	// InStockOne 添加单条库存记录
	InStockOne(*cmd.StockSaveCmd) (string, error)
	// QueryCategory 分类列表
	QueryCategory(query.CategoryPageQuery) ([]*dto.CategoryDto, error)
	// QueryStock 库存列表
	QueryStock(query.StockPageQuery) ([]*dto.StockDto, error)
	// TODO 批量添加记录、导入
}

// OrderSrv 订单
type OrderSrv interface {
	// CreateOrder 创建订单 返回付款url
	CreateOrder(*cmd.CreateOrderCmd) (string, error)
	// OrderCallBack 订单支付回调
	OrderCallBack(string) error
	// Cancel 取消订单
	Cancel(string) error
	// Get 获取订单详情
	Get(string) (*dto.OrderDto, error)
	// Query 获取订单列表
	Query(query.OrderPageQuery) ([]*dto.OrderListDto, error)
}

type SrvManager struct {
	AuthSrv      AuthSrv
	CommoditySrv CommoditySrv
	InventorySrv InventorySrv
	OrderSrv     OrderSrv
}

// NewAuthSrvManager wire
func NewAuthSrvManager(srv service.Service) *SrvManager {
	return &SrvManager{
		AuthSrv:      impl.NewAuthSrvImp(srv.UserSrv),
		CommoditySrv: impl.NewCommoditySrvImp(),
		InventorySrv: impl.NewInventorySrvImp(),
		OrderSrv:     impl.NewOrderSrvImp(),
	}
}
