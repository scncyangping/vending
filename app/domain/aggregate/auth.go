package aggregate

import (
	"vending/app/domain/entity"
)

type AuthAggregate struct {
	User  *entity.UserEntity
	Roles []*entity.RoleEntity `json:"roles"`
}

func (u *AuthAggregate) PwdCheck(pwd string) bool {
	return pwd == u.User.Pwd
}
