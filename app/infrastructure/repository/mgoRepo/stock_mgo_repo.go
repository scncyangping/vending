package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
)

var _ repo.StockRepo = (*StockMgoRepository)(nil)

type StockMgoRepository struct {
	mgo *mongo.MgoV
}

func NewStockMgoRepository(m *mongo.MgoV) *StockMgoRepository {
	return &StockMgoRepository{mgo: m}
}

func (s StockMgoRepository) SaveStock(entity *entity.StockEn) string {
	panic("implement me")
}

func (s StockMgoRepository) DeleteStock(s2 string) error {
	panic("implement me")
}

func (s StockMgoRepository) GetStockById(s2 string) *do.StockDo {
	panic("implement me")
}

func (s StockMgoRepository) ListStockBy(m map[string]interface{}) []*do.StockDo {
	panic("implement me")
}

func (s StockMgoRepository) ListStockPageBy(skip, limit int64, sort, filter interface{}) []*do.StockDo {
	panic("implement me")
}
