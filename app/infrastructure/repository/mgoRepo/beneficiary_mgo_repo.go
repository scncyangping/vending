package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
)

var _ repo.BeneficiaryRepo = (*BeneficiaryMgoRepository)(nil)

type BeneficiaryMgoRepository struct {
	mgo *mongo.MgoV
}

func NewBeneficiaryMgoRepository(m *mongo.MgoV) *BeneficiaryMgoRepository {
	return &BeneficiaryMgoRepository{mgo: m}
}

func (b BeneficiaryMgoRepository) SaveBeneficiary(entity *entity.BeneficiaryEn) string {
	panic("implement me")
}

func (b BeneficiaryMgoRepository) DeleteBeneficiary(s string) error {
	panic("implement me")
}

func (b BeneficiaryMgoRepository) GetBeneficiaryById(s string) *do.BeneficiaryDo {
	panic("implement me")
}

func (b BeneficiaryMgoRepository) ListBeneficiaryBy(m map[string]interface{}) []*do.BeneficiaryDo {
	panic("implement me")
}

func (b BeneficiaryMgoRepository) ListBeneficiaryPageBy(skip, limit int64, sort, filter interface{}) []*do.BeneficiaryDo {
	panic("implement me")
}
