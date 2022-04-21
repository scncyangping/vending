package repo

import (
	"vending/app/domain/entity"
	"vending/app/domain/vo"
)

type RoleRepo interface {
	SaveRole(entity *entity.RoleEntity) string
	DeleteRole(string) error
	GetRoleById(string) *vo.UserVo
	ListRoleBy(map[string]interface{}) []*vo.UserVo
	ListRolePageBy(skip, limit int64, sort, filter interface{}) []*vo.UserVo
}
