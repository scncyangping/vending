package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
)

var _ repo.PayDesRepo = (*PayDesMgoRepository)(nil)

type PayDesMgoRepository struct {
	mgo *mongo.MgoV
}

func NewPayDesMgoRepository(m *mongo.MgoV) *PayDesMgoRepository {
	return &PayDesMgoRepository{mgo: m}
}

func (p PayDesMgoRepository) SavePayDes(entity *entity.PayDesEn) string {
	panic("implement me")
}

func (p PayDesMgoRepository) DeletePayDes(s string) error {
	panic("implement me")
}

func (p PayDesMgoRepository) GetPayDesById(s string) *do.PayDesDo {
	panic("implement me")
}

func (p PayDesMgoRepository) ListPayDesBy(m map[string]interface{}) []*do.PayDesDo {
	panic("implement me")
}

func (p PayDesMgoRepository) ListPayDesPageBy(skip, limit int64, sort, filter interface{}) []*do.PayDesDo {
	panic("implement me")
}
