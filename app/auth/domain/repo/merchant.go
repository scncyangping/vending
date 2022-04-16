package repo

import (
	"context"
	"vending/app/auth/domain/entity"
)

type MerchantRepo interface {
	CreateMerChant(ctx context.Context, data *entity.Merchant) error
	UpdateMerChant(ctx context.Context, data *entity.Merchant) error
	RemoveMerChant(ctx context.Context, data *entity.Merchant) error
	QueryMerChant(ctx context.Context, repo MerChantSpecificationRepo) (data *entity.Merchant, err error)
	QueryMerChants(ctx context.Context, repo MerChantSpecificationRepo) (data []*entity.Merchant, err error)
}

type MerChantSpecificationRepo interface {
	ParameterCheck(ctx context.Context) error
	ToSql(ctx context.Context) interface{}
}
