package facotry

import "vending/app/ddd/domain/repo"

type AuthFactory struct {
	repo repo.JwtRepo
}
