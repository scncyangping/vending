package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
)

var _ repo.RoleRepo = (*RoleMgoRepository)(nil)

type RoleMgoRepository struct {
	mgo *mongo.MgoV
}

// NewRoleMgoRepository wire
func NewRoleMgoRepository() *RoleMgoRepository {
	return &RoleMgoRepository{
		mgo: mongo.OpCn("role"),
	}
}

func (r RoleMgoRepository) SaveRole(entity *entity.RoleEn) string {
	panic("implement me")
}

func (r RoleMgoRepository) DeleteRole(s string) error {
	panic("implement me")
}

func (r RoleMgoRepository) GetRoleById(s string) *do.RoleDo {
	panic("implement me")
}

func (r RoleMgoRepository) ListRoleBy(m map[string]interface{}) []*do.RoleDo {
	panic("implement me")
}

func (r RoleMgoRepository) ListRolePageBy(skip, limit int64, sort, filter interface{}) []*do.RoleDo {
	panic("implement me")
}
