package service

import (
	"vending/app/domain/dto"
	"vending/app/domain/repo"
)

type UserSrvImp struct {
	userRepo repo.UserRepo
}

func (u *UserSrvImp) CreateUser(rq dto.UserRegisterRq) error {
	return
}
