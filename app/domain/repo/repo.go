package repo

import (
	"vending/app/domain/entity"
	"vending/app/infrastructure/do"
)

type RoleRepo interface {
	SaveRole(entity *entity.RoleEn) string
	DeleteRole(string) error
	GetRoleById(string) *do.RoleDo
	ListRoleBy(map[string]interface{}) []*do.RoleDo
	ListRolePageBy(skip, limit int64, sort, filter interface{}) []*do.RoleDo
}

type CommodityRepo interface {
	SaveCommodity(entity *entity.CommodityEn, CategoryId string) string
	DeleteCommodity(string) error
	GetCommodityById(string) *do.CommodityDo
	ListCommodityBy(map[string]interface{}) []*do.CommodityDo
	ListCommodityPageBy(skip, limit int64, sort, filter interface{}) []*do.CommodityDo
}

type OrderRepo interface {
	SaveOrder(entity *entity.OrderEn) string
	DeleteOrder(string) error
	GetOrderById(string) *do.OrderDo
	ListOrderBy(map[string]interface{}) []*do.OrderDo
	ListOrderPageBy(skip, limit int64, sort, filter interface{}) []*do.OrderDo
}

type StockRepo interface {
	SaveStock(entity *entity.StockEn) string
	DeleteStock(string) error
	GetStockById(string) *do.StockDo
	ListStockBy(map[string]interface{}) []*do.StockDo
	ListStockPageBy(skip, limit int64, sort, filter interface{}) []*do.StockDo
}

type CategoryRepo interface {
	SaveCategory(entity *entity.CategoryEn) string
	DeleteCategory(string) error
	GetCategoryById(string) *do.CategoryDo
	ListCategoryBy(map[string]interface{}) []*do.CategoryDo
	ListCategoryPageBy(skip, limit int64, sort, filter interface{}) []*do.CategoryDo
}

type BeneficiaryRepo interface {
	SaveBeneficiary(entity *entity.BeneficiaryEn) string
	DeleteBeneficiary(string) error
	GetBeneficiaryById(string) *do.BeneficiaryDo
	ListBeneficiaryBy(map[string]interface{}) []*do.BeneficiaryDo
	ListBeneficiaryPageBy(skip, limit int64, sort, filter interface{}) []*do.BeneficiaryDo
}

type PayDesRepo interface {
	SavePayDes(entity *entity.PayDesEn) string
	DeletePayDes(string) error
	GetPayDesById(string) *do.PayDesDo
	ListPayDesBy(map[string]interface{}) []*do.PayDesDo
	ListPayDesPageBy(skip, limit int64, sort, filter interface{}) []*do.PayDesDo
}
