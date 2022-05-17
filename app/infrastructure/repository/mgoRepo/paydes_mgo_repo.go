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

var _ repo.PayDesRepo = (*PayDesMgoRepository)(nil)

type PayDesMgoRepository struct {
	mgo *mongo.MgoV
}

func (p *PayDesMgoRepository) SavePayDes(entity *entity.PayDesEn) (string, error) {
	var (
		do *do.PayDesDo
	)
	util.StructCopy(do, entity)
	do.CreateTime = util.NowTimestamp()
	do.UpdateTime = util.NowTimestamp()
	return p.mgo.InsertOne(do)
}

func (p *PayDesMgoRepository) UpdatePayDes(filter types.B, update types.B) error {
	if _, err := p.mgo.Update(filter, update); err != nil {
		log.Logger().Error("UpdatePayDes Error, %v", err)
		return err
	}
	return nil
}

func (p *PayDesMgoRepository) DeletePayDes(s string) error {
	if _, err := p.mgo.UpdateOne(types.B{"_id": s}, types.B{"isDeleted": 1}); err != nil {
		return err
	}
	return nil
}

func (p *PayDesMgoRepository) GetPayDesById(s string) (*do.PayDesDo, error) {
	var (
		err error
		do  do.PayDesDo
	)
	if err = p.mgo.FindOne(types.B{"_id": s}, &do); err != nil {
		log.Logger().Error("GetPayDesById Error, %v", err)
		return nil, err
	}
	return &do, nil
}

func (p *PayDesMgoRepository) ListPayDesBy(m types.B) ([]*do.PayDesDo, error) {
	var (
		err error
		dos []*do.PayDesDo
	)
	if err = p.mgo.Find(m, &dos); err != nil {
		log.Logger().Error("ListPayDesBy Error, %v", m)
		return nil, err
	}
	return dos, nil
}

func (p *PayDesMgoRepository) ListPayDesPageBy(skip, limit int64, sort, filter interface{}) ([]*do.PayDesDo, error) {
	var (
		err error
		dos []*do.PayDesDo
	)
	if err = p.mgo.FindBy(skip, limit, sort, filter, &dos); err != nil {
		log.Logger().Error("ListPayDesPageBy Error, %v", err)
		return nil, err
	}
	return dos, nil
}

func NewPayDesMgoRepository(m *mongo.MgoV) *PayDesMgoRepository {
	return &PayDesMgoRepository{mgo: m}
}
