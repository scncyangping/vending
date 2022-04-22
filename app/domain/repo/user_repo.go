package repo

import (
	"vending/app/domain/entity"
	"vending/app/infrastructure/do"
)

type UserRepo interface {
	SaveUser(entity *entity.UserEn) (string, error)
	DeleteUser(string) error
	GetUserById(string) (*do.UserDo, error)
	GetUserByName(string) (*do.UserDo, error)
	ListUserBy(map[string]interface{}) ([]*do.UserDo, error)
	ListUserPageBy(skip, limit int64, sort, filter interface{}) ([]*do.UserDo, error)
}
