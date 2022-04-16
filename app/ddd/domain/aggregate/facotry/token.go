package facotry

import "vending/app/ddd/domain/repo"

type TokenFactory struct {
	repo repo.JwtRepo
}

func (t *TokenFactory) NewToken() {
	t.repo.JwtCreate()
}
