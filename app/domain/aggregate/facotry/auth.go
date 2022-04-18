package facotry

import "vending/app/domain/repo"

type AuthFactory struct {
	repo repo.JwtRepo
}
