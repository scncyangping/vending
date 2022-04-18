package repo

import (
	"vending/app/ddd/domain/dto"
	"vending/app/ddd/infrastructure/pkg/tool"
)

// AuthServiceRepo 接口定义
type AuthServiceRepo interface {
}

type JwtRepo interface {
	JwtCreate(*dto.CreateTokenReq) (string, error)
	JwtCheck(string) (*tool.Claims, error)
}
