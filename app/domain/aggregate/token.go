package aggregate

import (
	"vending/app/ddd/domain/dto"
	"vending/app/ddd/domain/entity"
	"vending/app/ddd/domain/obj"
	"vending/app/ddd/domain/repo"
	"vending/app/ddd/infrastructure/pkg/tool"
)

type Token struct {
	Id       entity.Id
	JwtToken *obj.JwtToken
	Repo     repo.JwtRepo
}

func (t *Token) JwtCreate(o *dto.CreateTokenReq) (string, error) {
	t.JwtToken = &obj.JwtToken{
		Username: o.Name,
	}
	return tool.GenerateToken(t.JwtToken.Username)
}

func (t *Token) JwtCheck(s string) (*tool.Claims, error) {
	return tool.ParseToken(s)
}
