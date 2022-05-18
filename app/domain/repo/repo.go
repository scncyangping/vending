package repo

import (
	"vending/app/domain/entity"
	"vending/app/infrastructure/do"
	"vending/app/types"
)

type UserRepo interface {
	SaveUser(entity *entity.UserEn) (string, error)
	UpdateUser(queryFilter types.B, updateParm types.B) error
	DeleteUser(string) error
	GetUserById(string) (*do.UserDo, error)
	GetUserByName(string) (*do.UserDo, error)
	ListUserBy(map[string]interface{}) ([]*do.UserDo, error)
	ListUserPageBy(skip, limit int64, sort, filter interface{}) ([]*do.UserDo, error)
}

type RoleRepo interface {
	SaveRole(entity *entity.RoleEn) (string, error)
	UpdateRole(queryFilter types.B, updateParm types.B) error
	DeleteRole(string) error
	GetRoleById(string) (*do.RoleDo, error)
	ListRoleBy(types.B) ([]*do.RoleDo, error)
	ListRolePageBy(skip, limit int64, sort, filter interface{}) ([]*do.RoleDo, error)
}

type CommodityRepo interface {
	SaveCommodity(entity *entity.CommodityEn, CategoryId string) (string, error)
	UpdateCommodity(queryFilter types.B, updateParm types.B) error
	DeleteCommodity(string) error
	GetCommodityById(string) (*do.CommodityDo, error)
	ListCommodityBy(types.B) ([]*do.CommodityDo, error)
	ListCommodityPageBy(skip, limit int64, sort, filter interface{}) ([]*do.CommodityDo, error)
}

type OrderRepo interface {
	SaveOrder(entity *entity.OrderEn) (string, error)
	UpdateOrder(queryFilter types.B, updateParm types.B) error
	DeleteOrder(string) error
	GetOrderById(string) (*do.OrderDo, error)
	ListOrderBy(types.B) ([]*do.OrderDo, error)
	ListOrderPageBy(skip, limit int64, sort, filter interface{}) ([]*do.OrderDo, error)
}

type OrderTempRepo interface {
	SaveOrder(entity *entity.OrderEn) (string, error)
	UpdateOrder(queryFilter types.B, updateParm types.B) error
	DeleteOrder(string) error
	GetOrderById(string) (*do.OrderDo, error)
	ListOrderBy(types.B) ([]*do.OrderDo, error)
	ListOrderPageBy(skip, limit int64, sort, filter interface{}) ([]*do.OrderDo, error)
}

type StockRepo interface {
	SaveStock(entity *entity.StockEn) (string, error)
	UpdateStock(queryFilter types.B, updateParm types.B) (int64, error)
	DeleteStock(string) error
	GetStockById(string) (*do.StockDo, error)
	ListStockBy(types.B) ([]*do.StockDo, error)
	ListStockPageBy(skip, limit int64, sort, filter interface{}) ([]*do.StockDo, error)
}

type CategoryRepo interface {
	SaveCategory(*entity.CategoryEn) (string, error)
	UpdateCategory(queryFilter types.B, updateParm types.B) error
	DeleteCategoryByIds(ids []string) error
	DeleteCategory(string) error
	GetCategoryById(string) (*do.CategoryDo, error)
	GetCategoryByCategoryName(string) (*do.CategoryDo, error)
	ListCategoryBy(types.B) ([]*do.CategoryDo, error)
	ListCategoryPageBy(skip, limit int64, sort, filter interface{}) ([]*do.CategoryDo, error)
}

type BeneficiaryRepo interface {
	SaveBeneficiary(entity *entity.BeneficiaryEn) (string, error)
	UpdateBeneficiary(types.B, types.B) error
	DeleteBeneficiary(string) error
	GetBeneficiaryById(string) (*do.BeneficiaryDo, error)
	ListBeneficiaryBy(types.B) ([]*do.BeneficiaryDo, error)
	ListBeneficiaryPageBy(skip, limit int64, sort, filter interface{}) ([]*do.BeneficiaryDo, error)
}

type PayDesRepo interface {
	SavePayDes(entity *entity.PayDesEn) (string, error)
	UpdatePayDes(types.B, types.B) error
	DeletePayDes(string) error
	GetPayDesById(string) (*do.PayDesDo, error)
	ListPayDesBy(types.B) ([]*do.PayDesDo, error)
	ListPayDesPageBy(skip, limit int64, sort, filter interface{}) ([]*do.PayDesDo, error)
}
