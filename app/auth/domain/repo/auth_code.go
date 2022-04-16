package repo

import (
	"context"
	"vending/app/auth/domain/obj"
)

type AuthCodeRepo interface {
	CreateCode(ctx context.Context, data obj.CodeOpenId) error
	DelCode(ctx context.Context, repo AuthCodeSpecificationRepo) error
	QueryCode(ctx context.Context, repo AuthCodeSpecificationRepo) (data obj.CodeOpenId, err error)
}

type AuthCodeSpecificationRepo interface {
	ParameterCheck(ctx context.Context) error
	ToSql(ctx context.Context) interface{}
}
