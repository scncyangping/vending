package repo

import (
	"vending/app/domain/entity"
	"vending/app/domain/vo"
)

type UserRepo interface {
	SaveUser(entity *entity.UserEntity) string
	DeleteUser(string) error
	GetUserById(string) *vo.UserVo
	ListUserBy(map[string]interface{}) []*vo.UserVo
	ListUserPageBy(skip, limit int64, sort, filter interface{}) []*vo.UserVo
}
