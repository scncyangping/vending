package service

import "vending/app/domain/dto"

type UserService interface {
	Register(rq dto.UserRegisterRq) error
}
