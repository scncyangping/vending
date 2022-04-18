package facotry

import (
	"vending/app/ddd/domain/aggregate"
	"vending/app/ddd/domain/entity"
	"vending/app/ddd/domain/repo"
)

type TokenFactory struct {
	repo repo.JwtRepo
}

func (t *TokenFactory) NewToken() *aggregate.Token {
	return &aggregate.Token{
		Id:   entity.NewId(),
		Repo: t.repo,
	}
}
