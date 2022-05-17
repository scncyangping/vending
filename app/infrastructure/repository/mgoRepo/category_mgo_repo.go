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

var _ repo.CategoryRepo = (*CategoryMgoRepository)(nil)

type CategoryMgoRepository struct {
	mgo *mongo.MgoV
}

func NewCategoryMgoRepository(m *mongo.MgoV) *CategoryMgoRepository {
	return &CategoryMgoRepository{mgo: m}
}

func (c *CategoryMgoRepository) SaveCategory(entity *entity.CategoryEn) (string, error) {
	var (
		CategoryDo *do.CategoryDo
	)
	util.StructCopy(CategoryDo, entity)
	CategoryDo.CreateTime = util.NowTimestamp()
	CategoryDo.UpdateTime = util.NowTimestamp()
	return c.mgo.InsertOne(CategoryDo)
}

func (c *CategoryMgoRepository) UpdateCategory(filter types.B, update types.B) error {
	if _, err := c.mgo.Update(filter, update); err != nil {
		log.Logger().Error("UpdateCategory Error, %v", err)
		return err
	}
	return nil
}

func (c *CategoryMgoRepository) DeleteCategory(s string) error {
	if _, err := c.mgo.UpdateOne(types.B{"_id": s}, types.B{"isDeleted": 1}); err != nil {
		log.Logger().Error("DeleteCategory Error, %v", err)
		return err
	}
	return nil
}

func (c *CategoryMgoRepository) GetCategoryById(s string) (*do.CategoryDo, error) {
	var (
		err error
		cg  do.CategoryDo
	)
	if err = c.mgo.FindOne(types.B{"_id": s}, &cg); err != nil {
		log.Logger().Error("GetCategoryById Error, %v", err)
		return nil, err
	}
	return &cg, nil
}

func (c *CategoryMgoRepository) ListCategoryBy(m map[string]interface{}) ([]*do.CategoryDo, error) {
	var (
		err error
		cgs []*do.CategoryDo
	)
	if err = c.mgo.Find(m, &cgs); err != nil {
		log.Logger().Error("ListCategoryBy Error, %v", m)
		return nil, err
	}
	return cgs, nil
}

func (c *CategoryMgoRepository) ListCategoryPageBy(skip, limit int64, sort, filter interface{}) ([]*do.CategoryDo, error) {
	var (
		err error
		cgs []*do.CategoryDo
	)
	if err = c.mgo.FindBy(skip, limit, sort, filter, &cgs); err != nil {
		log.Logger().Error("ListCategoryPageBy Error, %v", err)
		return nil, err
	}
	return cgs, nil
}
