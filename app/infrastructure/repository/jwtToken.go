package repository

import (
	"vending/app/domain/dto"
	"vending/app/domain/repo"
	"vending/app/infrastructure/pkg/tool"
)

var _ repo.JwtRepo = (*AuthToken)(nil)

type AuthToken struct {
	repository
}

func (a AuthToken) JwtCreate(req *dto.CreateTokenReq) (string, error) {
	return tool.GenerateToken(req.Name)
}

func (a AuthToken) JwtCheck(s string) (*tool.Claims, error) {
	return tool.ParseToken(s)
}
