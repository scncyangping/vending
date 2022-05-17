package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
)

var _ repo.CommodityRepo = (*CommodityMgoRepository)(nil)

type CommodityMgoRepository struct {
	mgo *mongo.MgoV
}

func NewCommodityMgoRepository(m *mongo.MgoV) *CommodityMgoRepository {
	return &CommodityMgoRepository{mgo: m}
}

func (c CommodityMgoRepository) SaveCommodity(entity *entity.CommodityEn, CategoryId string) string {
	panic("implement me")
}

func (c CommodityMgoRepository) DeleteCommodity(s string) error {
	panic("implement me")
}

func (c CommodityMgoRepository) GetCommodityById(s string) *do.CommodityDo {
	panic("implement me")
}

func (c CommodityMgoRepository) ListCommodityBy(m map[string]interface{}) []*do.CommodityDo {
	panic("implement me")
}

func (c CommodityMgoRepository) ListCommodityPageBy(skip, limit int64, sort, filter interface{}) []*do.CommodityDo {
	panic("implement me")
}
