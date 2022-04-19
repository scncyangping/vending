package auth

import (
	"vending/app/domain/dto"
	"vending/app/domain/obj"
	"vending/app/infrastructure/pkg/tool"
)

type JwtTokenAuth struct {
}

func NewJwtTokenAuth() JwtTokenAuth {
	return JwtTokenAuth{}
}

func (j JwtTokenAuth) GenerateToken(re *dto.JwtAuthTokenRe) (string, error) {
	return tool.GenerateToken(obj.JwtToken{
		Username: re.UserName,
	})
}

func (j JwtTokenAuth) ParseToken(s string) (*obj.Claims, error) {
	return tool.ParseToken(s)
}
