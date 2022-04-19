package service

import (
	"vending/app/domain/dto"
	"vending/app/domain/obj"
)

// AuthSrv Auth相关应单独抽离出来
type AuthSrv interface {
	GenerateToken(re *dto.JwtAuthTokenRe) (string, error)
	ParseToken(string) (*obj.Claims, error)
}
