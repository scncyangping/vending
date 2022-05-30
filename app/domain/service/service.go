package service

import (
	"vending/app/application/cqe/cmd"
	"vending/app/domain/entity"
	"vending/app/domain/obj"
	"vending/app/domain/service/imp/auth"
	"vending/app/domain/service/imp/commodity"
	"vending/app/domain/service/imp/inventory"
	"vending/app/domain/service/imp/order"
	"vending/app/domain/service/imp/pay"
	"vending/app/domain/vo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/repository"
	"vending/app/types"
)

type AuthService interface {
	LoginByName(name, pwd string) (*vo.UserVo, error)
	Register(*entity.UserEn) (string, error)
}

//type CategoryService interface {
//	SaveCategory(req *dto.CategorySaveReq) (string, error)
//}

type DoCommodityService interface {
	QueryCommoditiesByIds(ids []string) ([]*entity.CommodityEn, error)
	SaveCommodity(req *cmd.CommoditySaveCmd) (string, error)
	DeleteCommodityBatch(s []string) error
	QueryCommodityPageBy(skip, limit int64, sort, filter any) ([]*do.CommodityDo, error)
}

type DoInventoryService interface {
	SaveCategory(req *cmd.CategorySaveCmd) (string, error)
	RemoveCategoryByIds(ids []string) error
	QueryCategoryPageBy(skip, limit int64, sort, filter any) ([]*do.CategoryDo, error)
	QueryStockPageBy(skip, limit int64, sort, filter any) ([]*do.StockDo, error)
}

type DoOrderService interface {
	CreateTempOrderOne(items []obj.OrderItemObj, desObj obj.PayDesObj) (string, float64, error)
	SaveOrder(orderId string) error
	GetOrderById(id string) (*do.OrderDo, error)
	GetTempOrderById(id string) (*do.OrderDo, error)
	QueryOrderPageBy(skip, limit int64, sort, filter any) ([]*do.OrderDo, error)
	QueryTempOrderPageBy(skip, limit int64, sort, filter any) ([]*do.OrderDo, error)
}

type DoPayService interface {
	SaveBeneficiary(beneficiaryType types.BeneficiaryType, data any, userId string) (string, error)
	PayUrl(orderId string, amount float64) (string, error)
}
type Service struct {
	UserSrv        *auth.UserServiceImpl
	DoCommoditySrv *commodity.DoCommoditySrvImpl
	DoInventorySrv *inventory.DoInventorySrvImpl
	DoOrderSrv     *order.DoOrderSrvImpl
	DoPayService   *pay.DoPaySrvImpl
}

// NewService wire
func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserSrv:        auth.NewUserServiceImpl(repo.UserRepo),
		DoCommoditySrv: commodity.NewDoCommoditySrvImpl(repo.CommodityRepo, repo.CategoryRepo),
		DoInventorySrv: inventory.NewDoInventorySrvImpl(repo.CategoryRepo, repo.StockRepo),
		DoOrderSrv:     order.NewDoOrderSrvImpl(repo.BeneficiaryRepo, repo.OrderRepo, repo.OrderTempRepo),
		DoPayService:   pay.NewDoPaySrvImpl(repo.BeneficiaryRepo),
	}
}
