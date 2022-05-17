package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
)

var _ repo.CategoryRepo = (*CategoryMgoRepository)(nil)

type CategoryMgoRepository struct {
	mgo *mongo.MgoV
}

func NewCategoryMgoRepository(m *mongo.MgoV) *CategoryMgoRepository {
	return &CategoryMgoRepository{mgo: m}
}

func (c CategoryMgoRepository) SaveCategory(entity *entity.CategoryEn) string {
	panic("implement me")
}

func (c CategoryMgoRepository) DeleteCategory(s string) error {
	panic("implement me")
}

func (c CategoryMgoRepository) GetCategoryById(s string) *do.CategoryDo {
	panic("implement me")
}

func (c CategoryMgoRepository) ListCategoryBy(m map[string]interface{}) []*do.CategoryDo {
	panic("implement me")
}

func (c CategoryMgoRepository) ListCategoryPageBy(skip, limit int64, sort, filter interface{}) []*do.CategoryDo {
	panic("implement me")
}
