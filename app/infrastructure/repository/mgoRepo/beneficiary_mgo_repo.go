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

var _ repo.BeneficiaryRepo = (*BeneficiaryMgoRepository)(nil)

type BeneficiaryMgoRepository struct {
	mgo *mongo.MgoV
}

func NewBeneficiaryMgoRepository(m *mongo.MgoV) *BeneficiaryMgoRepository {
	return &BeneficiaryMgoRepository{mgo: m}
}

func (b BeneficiaryMgoRepository) SaveBeneficiary(entity *entity.BeneficiaryEn) (string, error) {
	var (
		beneficiaryDo *do.BeneficiaryDo
	)
	util.StructCopy(beneficiaryDo, entity)
	beneficiaryDo.CreateTime = util.NowTimestamp()
	beneficiaryDo.UpdateTime = util.NowTimestamp()
	return b.mgo.InsertOne(beneficiaryDo)
}

func (b BeneficiaryMgoRepository) DeleteBeneficiary(s string) error {
	b.mgo.UpdateOne(types.B{"_id": s}, types.B{"isDeleted": 1})
	return nil
}

func (b BeneficiaryMgoRepository) GetBeneficiaryById(s string) (*do.BeneficiaryDo, error) {
	var (
		err error
		bfa do.BeneficiaryDo
	)
	if err = b.mgo.FindOne(types.B{"_id": s}, &bfa); err != nil {
		log.Logger().Error("GetBeneficiaryById Error, %v", err)
		return nil, err
	}
	return &bfa, nil
}

func (b BeneficiaryMgoRepository) ListBeneficiaryBy(m map[string]interface{}) ([]*do.BeneficiaryDo, error) {
	var (
		err error
		bfs []*do.BeneficiaryDo
	)
	if err = b.mgo.Find(m, &bfs); err != nil {
		log.Logger().Error("ListBeneficiaryBy Error, %v", m)
		return nil, err
	}
	return bfs, nil
}

func (b BeneficiaryMgoRepository) ListBeneficiaryPageBy(skip, limit int64, sort, filter interface{}) ([]*do.BeneficiaryDo, error) {
	var (
		err error
		bfs []*do.BeneficiaryDo
	)
	if err = b.mgo.FindBy(skip, limit, sort, filter, &bfs); err != nil {
		log.Logger().Error("ListBeneficiaryPageBy Error, %v", err)
		return nil, err
	}
	return bfs, nil
}
