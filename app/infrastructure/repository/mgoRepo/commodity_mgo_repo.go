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

var _ repo.CommodityRepo = (*CommodityMgoRepository)(nil)

type CommodityMgoRepository struct {
	mgo *mongo.MgoV
}

func (c *CommodityMgoRepository) SaveCommodity(entity *entity.CommodityEn, CategoryId string) (string, error) {
	var (
		commodityDo *do.CommodityDo
	)
	util.StructCopy(commodityDo, entity)
	commodityDo.CreateTime = util.NowTimestamp()
	commodityDo.UpdateTime = util.NowTimestamp()
	commodityDo.CategoryId = CategoryId
	return c.mgo.InsertOne(commodityDo)
}

func (c *CommodityMgoRepository) UpdateCommodity(filter types.B, update types.B) error {
	if _, err := c.mgo.Update(filter, update); err != nil {
		log.Logger().Error("UpdateCommodity Error, %v", err)
		return err
	}
	return nil
}

func (c *CommodityMgoRepository) DeleteCommodity(s string) error {
	if _, err := c.mgo.UpdateOne(types.B{"_id": s}, types.B{"isDeleted": 1}); err != nil {
		return err
	}
	return nil
}

func (c *CommodityMgoRepository) DeleteCommodityBatch(s []string) error {
	if _, err := c.mgo.UpdateOne(types.B{"_id": types.B{"$in": s}}, types.B{"isDeleted": 1}); err != nil {
		return err
	}
	return nil
}

func (c *CommodityMgoRepository) GetCommodityById(s string) (*do.CommodityDo, error) {
	var (
		err error
		do  do.CommodityDo
	)
	if err = c.mgo.FindOne(types.B{"_id": s}, &do); err != nil {
		log.Logger().Error("GetCommodityById Error, %v", err)
		return nil, err
	}
	return &do, nil
}

func (c *CommodityMgoRepository) ListCommodityBy(m types.B) ([]*do.CommodityDo, error) {
	var (
		err error
		dos []*do.CommodityDo
	)
	if err = c.mgo.Find(m, &dos); err != nil {
		log.Logger().Error("ListCommodityBy Error, %v", m)
		return nil, err
	}
	return dos, nil
}

func (c *CommodityMgoRepository) ListCommodityPageBy(skip, limit int64, sort, filter any) ([]*do.CommodityDo, error) {
	var (
		err error
		dos []*do.CommodityDo
	)
	if err = c.mgo.FindBy(skip, limit, sort, filter, &dos); err != nil {
		log.Logger().Error("ListCommodityPageBy Error, %v", err)
		return nil, err
	}
	return dos, nil
}

func NewCommodityMgoRepository(m *mongo.MgoV) *CommodityMgoRepository {
	return &CommodityMgoRepository{mgo: m}
}
