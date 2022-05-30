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

var _ repo.StockRepo = (*StockMgoRepository)(nil)

type StockMgoRepository struct {
	mgo *mongo.MgoV
}

func (s *StockMgoRepository) SaveStock(entity *entity.StockEn) (string, error) {
	var (
		do *do.StockDo
	)
	util.StructCopy(do, entity)
	do.CreateTime = util.NowTimestamp()
	do.UpdateTime = util.NowTimestamp()
	return s.mgo.InsertOne(do)
}

func (s *StockMgoRepository) UpdateStock(filter any, update any) (int64, error) {
	if count, err := s.mgo.Update(filter, update); err != nil {
		log.Logger().Error("UpdateStock Error, %v", err)
		return count, err
	}
	return 0, nil
}

func (s *StockMgoRepository) DeleteStock(id string) error {
	if _, err := s.mgo.UpdateOne(types.B{"_id": id}, types.B{"isDeleted": 1}); err != nil {
		return err
	}
	return nil
}

func (s *StockMgoRepository) GetStockById(id string) (*do.StockDo, error) {
	var (
		err error
		do  do.StockDo
	)
	if err = s.mgo.FindOne(types.B{"_id": id}, &do); err != nil {
		log.Logger().Error("GetStockById Error, %v", err)
		return nil, err
	}
	return &do, nil
}

func (s *StockMgoRepository) ListStockBy(m any) ([]*do.StockDo, error) {
	var (
		err error
		dos []*do.StockDo
	)
	if err = s.mgo.Find(m, &dos); err != nil {
		log.Logger().Error("ListStockBy Error, %v", m)
		return nil, err
	}
	return dos, nil
}

func (s *StockMgoRepository) ListStockByIdsAndStatus(ids []string, status types.StockStatus) ([]*do.StockDo, error) {
	if l, err := s.ListStockBy(types.B{"categoryId": types.B{"$in": ids}, "Status": status}); err != nil {
		return l, err
	} else {
		return l, nil
	}
}

func (s *StockMgoRepository) ListStockPageBy(skip, limit int64, sort, filter any) ([]*do.StockDo, error) {
	var (
		err error
		dos []*do.StockDo
	)
	if err = s.mgo.FindBy(skip, limit, sort, filter, &dos); err != nil {
		log.Logger().Error("ListStockPageBy Error, %v", err)
		return nil, err
	}
	return dos, nil
}

func NewStockMgoRepository(m *mongo.MgoV) *StockMgoRepository {
	return &StockMgoRepository{mgo: m}
}
