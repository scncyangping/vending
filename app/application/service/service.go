package service

import "vending/app/domain/dto"

type AuthService interface {
	AuthByJWT(dto.JwtAuthRe) dto.JwtAuthRp
	Register(rq dto.UserRegisterRq) (error, string)
}
