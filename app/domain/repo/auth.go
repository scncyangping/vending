package repo

import (
	"vending/app/domain/dto"
	"vending/app/infrastructure/pkg/tool"
)

// AuthServiceRepo 接口定义
type AuthServiceRepo interface {
	JwtRepo
}

type JwtRepo interface {
	JwtCreate(*dto.CreateTokenReq) (string, error)
	JwtCheck(string) (*tool.Claims, error)
}
