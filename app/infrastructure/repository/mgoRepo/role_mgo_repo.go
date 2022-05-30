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

var _ repo.RoleRepo = (*RoleMgoRepository)(nil)

type RoleMgoRepository struct {
	mgo *mongo.MgoV
}

func (r *RoleMgoRepository) SaveRole(entity *entity.RoleEn) (string, error) {
	var (
		do *do.RoleDo
	)
	util.StructCopy(do, entity)
	do.CreateTime = util.NowTimestamp()
	do.UpdateTime = util.NowTimestamp()
	return r.mgo.InsertOne(do)
}

func (r *RoleMgoRepository) UpdateRole(filter any, update any) error {
	if _, err := r.mgo.Update(filter, update); err != nil {
		log.Logger().Error("UpdateRole Error, %v", err)
		return err
	}
	return nil
}

func (r *RoleMgoRepository) DeleteRole(s string) error {
	if _, err := r.mgo.UpdateOne(types.B{"_id": s}, types.B{"isDeleted": 1}); err != nil {
		return err
	}
	return nil
}

func (r *RoleMgoRepository) GetRoleById(s string) (*do.RoleDo, error) {
	var (
		err error
		do  do.RoleDo
	)
	if err = r.mgo.FindOne(types.B{"_id": s}, &do); err != nil {
		log.Logger().Error("GetRoleById Error, %v", err)
		return nil, err
	}
	return &do, nil
}

func (r *RoleMgoRepository) ListRoleBy(m any) ([]*do.RoleDo, error) {
	var (
		err error
		dos []*do.RoleDo
	)
	if err = r.mgo.Find(m, &dos); err != nil {
		log.Logger().Error("ListRoleBy Error, %v", m)
		return nil, err
	}
	return dos, nil
}

func (r *RoleMgoRepository) ListRolePageBy(skip, limit int64, sort, filter any) ([]*do.RoleDo, error) {
	var (
		err error
		dos []*do.RoleDo
	)
	if err = r.mgo.FindBy(skip, limit, sort, filter, &dos); err != nil {
		log.Logger().Error("ListRolePageBy Error, %v", err)
		return nil, err
	}
	return dos, nil
}

func NewRoleMgoRepository(m *mongo.MgoV) *RoleMgoRepository {
	return &RoleMgoRepository{mgo: m}
}
