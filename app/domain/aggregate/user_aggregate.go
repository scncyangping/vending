package aggregate

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
)

type UserAggregate struct {
	User     *entity.UserEntity
	Roles    []string
	userRepo *repo.UserRepo
}

func (u *UserAggregate) Save() {

}
