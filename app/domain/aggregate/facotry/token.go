package facotry

import (
	"vending/app/domain/aggregate"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
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
