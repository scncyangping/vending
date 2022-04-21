package service

import (
	"vending/app/domain/aggregate/facotry"
	"vending/app/domain/dto"
)

type AuthSrvImp struct {
	*facotry.AuthFactory
}

func (a *AuthSrvImp) AuthByJWT(re dto.JwtAuthRe) dto.JwtAuthRp {
	panic("implement me")
}

func (a *AuthSrvImp) Register(rq dto.UserRegisterRq) (error, string) {
	panic("implement me")
}
