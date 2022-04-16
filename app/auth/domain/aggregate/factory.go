package aggregate

import (
	"vending/app/auth/infrastructure/repository"
)

type Factory struct {
	*AuthFactory
}

func NewFactory(repo *repository.Repository) *Factory {
	return &Factory{
		&AuthFactory{
			merchantRepo:  repo.Merchant,
			authCodeRepo:  repo.AuthCode,
			authTokenRepo: repo.AuthToken,
		},
	}
}
