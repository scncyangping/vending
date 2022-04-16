package repo

import (
	"vending/app/ddd/domain/obj"
)

// AuthServiceRepo 接口定义
type AuthServiceRepo interface {
}

type JwtRepo interface {
	JwtCreate(*obj.JwtToken) (string, error)
	JwtCheck(string) (*obj.JwtToken, error)
}
