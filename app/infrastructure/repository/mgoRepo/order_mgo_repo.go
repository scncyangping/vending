package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
)

var _ repo.OrderRepo = (*OrderMgoRepository)(nil)

type OrderMgoRepository struct {
	mgo *mongo.MgoV
}

func NewOrderMgoRepository(m *mongo.MgoV) *OrderMgoRepository {
	return &OrderMgoRepository{mgo: m}
}
func (o OrderMgoRepository) SaveOrder(entity *entity.OrderEn) string {
	panic("implement me")
}

func (o OrderMgoRepository) DeleteOrder(s string) error {
	panic("implement me")
}

func (o OrderMgoRepository) GetOrderById(s string) *do.OrderDo {
	panic("implement me")
}

func (o OrderMgoRepository) ListOrderBy(m map[string]interface{}) []*do.OrderDo {
	panic("implement me")
}

func (o OrderMgoRepository) ListOrderPageBy(skip, limit int64, sort, filter interface{}) []*do.OrderDo {
	panic("implement me")
}
