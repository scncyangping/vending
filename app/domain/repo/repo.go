package repo

import (
	"vending/app/domain/entity"
	"vending/app/infrastructure/do"
	"vending/app/types"
)

type UserRepo interface {
	SaveUser(entity *entity.UserEn) (string, error)
	UpdateUser(queryFilter any, updateParm any) error
	DeleteUser(string) error
	GetUserById(string) (*do.UserDo, error)
	GetUserByName(string) (*do.UserDo, error)
	ListUserBy(any) ([]*do.UserDo, error)
	ListUserPageBy(skip, limit int64, sort, filter any) ([]*do.UserDo, error)
}

type RoleRepo interface {
	SaveRole(entity *entity.RoleEn) (string, error)
	UpdateRole(queryFilter any, updateParm any) error
	DeleteRole(string) error
	GetRoleById(string) (*do.RoleDo, error)
	ListRoleBy(any) ([]*do.RoleDo, error)
	ListRolePageBy(skip, limit int64, sort, filter any) ([]*do.RoleDo, error)
}

type CommodityRepo interface {
	SaveCommodity(entity *entity.CommodityEn, CategoryId string) (string, error)
	UpdateCommodity(any, any) error
	DeleteCommodity(string) error
	DeleteCommodityBatch(s []string) error
	GetCommodityById(string) (*do.CommodityDo, error)
	ListCommodityByIds(ids []string) ([]*do.CommodityDo, error)
	ListCommodityPageBy(skip, limit int64, sort, filter any) ([]*do.CommodityDo, error)
}

type OrderRepo interface {
	SaveOrder(entity *entity.OrderEn) (string, error)
	UpdateOrder(any, any) error
	DeleteOrder(string) error
	GetOrderById(string) (*do.OrderDo, error)
	ListOrderBy(any) ([]*do.OrderDo, error)
	ListOrderPageBy(skip, limit int64, sort, filter any) ([]*do.OrderDo, error)
}

type OrderTempRepo interface {
	SaveOrder(entity *entity.OrderEn) (string, error)
	UpdateOrder(any, any) error
	DeleteOrder(string) error
	GetOrderById(string) (*do.OrderDo, error)
	ListOrderBy(any) ([]*do.OrderDo, error)
	ListOrderPageBy(skip, limit int64, sort, filter any) ([]*do.OrderDo, error)
}

type StockRepo interface {
	SaveStock(entity *entity.StockEn) (string, error)
	UpdateStock(queryFilter any, updateParm any) (int64, error)
	DeleteStock(string) error
	GetStockById(string) (*do.StockDo, error)
	ListStockBy(any) ([]*do.StockDo, error)
	ListStockByIdsAndStatus(ids []string, status types.StockStatus) ([]*do.StockDo, error)
	ListStockPageBy(skip, limit int64, sort, filter any) ([]*do.StockDo, error)
}

type CategoryRepo interface {
	SaveCategory(*entity.CategoryEn) (string, error)
	UpdateCategory(any, any) error
	DeleteCategoryByIds(ids []string) error
	DeleteCategory(string) error
	GetCategoryById(string) (*do.CategoryDo, error)
	GetCategoryByCategoryName(string) (*do.CategoryDo, error)
	ListCategoryBy(any) ([]*do.CategoryDo, error)
	ListCategoryPageBy(skip, limit int64, sort, filter any) ([]*do.CategoryDo, error)
}

type BeneficiaryRepo interface {
	SaveBeneficiary(entity *entity.BeneficiaryEn) (string, error)
	UpdateBeneficiary(any, any) error
	DeleteBeneficiary(string) error
	GetBeneficiaryById(string) (*do.BeneficiaryDo, error)
	GetBeneficiaryByOwnerIdAndType(string, types.BeneficiaryType) (*do.BeneficiaryDo, error)
	GetBeneficiaryByOwnerIdOrTypeDefault(s string, beneficiaryType types.BeneficiaryType) (*do.BeneficiaryDo, error) // 根据ownerId及type查询，若对应type不存在,返回任意其他类型支付方式
	ListBeneficiaryBy(any) ([]*do.BeneficiaryDo, error)
	ListBeneficiaryPageBy(skip, limit int64, sort, filter any) ([]*do.BeneficiaryDo, error)
}

//type PayDesRepo interface {
//	SavePayDes(entity *entity.PayDes) (string, error)
//	UpdatePayDes(any, any) error
//	DeletePayDes(string) error
//	GetPayDesById(string) (*do.PayDesDo, error)
//	ListPayDesBy(any) ([]*do.PayDesDo, error)
//	ListPayDesPageBy(skip, limit int64, sort, filter any) ([]*do.PayDesDo, error)
//}
