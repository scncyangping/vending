package repo

import (
	"vending/app/domain/entity"
	"vending/app/domain/vo"
)

type RoleRepo interface {
	SaveRole(entity *entity.RoleEn) string
	DeleteRole(string) error
	GetRoleById(string) *vo.RoleVo
	ListRoleBy(map[string]interface{}) []*vo.RoleVo
	ListRolePageBy(skip, limit int64, sort, filter interface{}) []*vo.RoleVo
}
