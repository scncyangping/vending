package aggregate

import (
	"vending/app/domain/entity"
	"vending/app/domain/obj"
	"vending/app/domain/repo"
)

type Token struct {
	Id       entity.Id
	JwtToken *obj.JwtToken
	Repo     repo.JwtRepo
}

//func (t *Token) JwtCreate(o *dto.CreateTokenReq) (string, error) {
//	t.JwtToken = &obj.JwtToken{
//		Username: o.Name,
//	}
//	return tool.GenerateToken(t.JwtToken.Username)
//}
//
//func (t *Token) JwtCheck(s string) (*tool.Claims, error) {
//	return tool.ParseToken(s)
//}
