package repository

import (
	"vending/app/infrastructure/pkg/database/mongo"
	"vending/app/infrastructure/repository/mgoRepo"
)

type Repository struct {
	UserRepo        *mgoRepo.UserMgoRepository
	BeneficiaryRepo *mgoRepo.BeneficiaryMgoRepository
	CategoryRepo    *mgoRepo.CategoryMgoRepository
	CommodityRepo   *mgoRepo.CommodityMgoRepository
	OrderRepo       *mgoRepo.OrderMgoRepository
	OrderTempRepo   *mgoRepo.OrderMgoRepository
	StockRepo       *mgoRepo.StockMgoRepository
	RoleRepo        *mgoRepo.RoleMgoRepository
}

// NewRepository wire
func NewRepository() *Repository {
	return &Repository{
		UserRepo:        mgoRepo.NewUserRepository(mongo.OpCn("user")),
		BeneficiaryRepo: mgoRepo.NewBeneficiaryMgoRepository(mongo.OpCn("beneficiary")),
		CategoryRepo:    mgoRepo.NewCategoryMgoRepository(mongo.OpCn("category")),
		CommodityRepo:   mgoRepo.NewCommodityMgoRepository(mongo.OpCn("commodity")),
		OrderRepo:       mgoRepo.NewOrderMgoRepository(mongo.OpCn("order")),
		OrderTempRepo:   mgoRepo.NewOrderMgoRepository(mongo.OpCn("order_temp")),
		StockRepo:       mgoRepo.NewStockMgoRepository(mongo.OpCn("stock")),
		RoleRepo:        mgoRepo.NewRoleMgoRepository(mongo.OpCn("role")),
	}
}
