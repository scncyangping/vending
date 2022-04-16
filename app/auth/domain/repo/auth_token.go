package repo

import (
	"context"
	"vending/app/auth/domain/obj"
)

type AuthTokenRepo interface {
	CreateAuthToken(ctx context.Context, data obj.AuthToken) error
	UpdateAuthToken(ctx context.Context, data obj.AuthToken) error
	QueryAuthToken(ctx context.Context, repo AuthTokenSpecificationRepo) (obj.AuthToken, error)
}

type AuthTokenSpecificationRepo interface {
	ParameterCheck(ctx context.Context) error
	ToSql(ctx context.Context) interface{}
}
