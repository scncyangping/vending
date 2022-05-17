package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
	"vending/app/infrastructure/pkg/log"
	"vending/app/infrastructure/pkg/util"
	"vending/app/types"
)

var _ repo.OrderRepo = (*OrderMgoRepository)(nil)

type OrderMgoRepository struct {
	mgo *mongo.MgoV
}

func (o OrderMgoRepository) SaveOrder(entity *entity.OrderEn) (string, error) {
	var (
		do *do.OrderDo
	)
	util.StructCopy(do, entity)
	do.CreateTime = util.NowTimestamp()
	do.UpdateTime = util.NowTimestamp()
	return o.mgo.InsertOne(do)
}

func (o OrderMgoRepository) DeleteOrder(s string) error {
	if _, err := o.mgo.UpdateOne(types.B{"_id": s}, types.B{"isDeleted": 1}); err != nil {
		return err
	}
	return nil
}

func (o OrderMgoRepository) GetOrderById(s string) (*do.OrderDo, error) {
	var (
		err error
		do  do.OrderDo
	)
	if err = o.mgo.FindOne(types.B{"_id": s}, &do); err != nil {
		log.Logger().Error("GetOrderById Error, %v", err)
		return nil, err
	}
	return &do, nil
}

func (o OrderMgoRepository) ListOrderBy(m map[string]interface{}) ([]*do.OrderDo, error) {
	var (
		err error
		dos []*do.OrderDo
	)
	if err = o.mgo.Find(m, &dos); err != nil {
		log.Logger().Error("ListOrderBy Error, %v", m)
		return nil, err
	}
	return dos, nil
}

func (o OrderMgoRepository) ListOrderPageBy(skip, limit int64, sort, filter interface{}) ([]*do.OrderDo, error) {
	var (
		err error
		dos []*do.OrderDo
	)
	if err = o.mgo.FindBy(skip, limit, sort, filter, &dos); err != nil {
		log.Logger().Error("ListOrderPageBy Error, %v", err)
		return nil, err
	}
	return dos, nil
}

func NewOrderMgoRepository(m *mongo.MgoV) *OrderMgoRepository {
	return &OrderMgoRepository{mgo: m}
}
