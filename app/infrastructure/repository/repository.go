package repository

import "vending/app/infrastructure/repository/auth"

type Repository struct {
	UserRepo *auth.UserMgoRepository
}

// NewRepository wire
func NewRepository(u *auth.UserMgoRepository) *Repository {
	return &Repository{
		UserRepo: u,
	}
}
