package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"vending/app/auth/domain/entity"
	"vending/app/auth/domain/repo"
	"vending/app/auth/infrastructure/pkg/hcode"
	"vending/app/auth/infrastructure/pkg/log"
)

var _ repo.MerchantRepo = (*Merchant)(nil)

type Merchant struct {
	repository
}

func (m *Merchant) getMerChant() *mongo.Collection {
	return m.mgo.Database("gs").Collection("auth_merChant")
}

func (m *Merchant) CreateMerChant(ctx context.Context, data *entity.Merchant) error {
	if _, err := m.getMerChant().InsertOne(ctx, data); err != nil {
		log.GetLogger().Error("[repository] Merchant QueryMerChant FindOne", zap.Any("data", data), zap.Error(err))
		return hcode.MgoExecErr
	}
	return nil
}

func (m *Merchant) UpdateMerChant(ctx context.Context, data *entity.Merchant) error {
	return nil
}

func (m *Merchant) RemoveMerChant(ctx context.Context, data *entity.Merchant) error {
	return nil
}

func (m *Merchant) QueryMerChant(ctx context.Context, repo repo.MerChantSpecificationRepo) (data *entity.Merchant, err error) {
	data = new(entity.Merchant)
	if err = m.getMerChant().FindOne(ctx, repo.ToSql(ctx)).Decode(data); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, hcode.ResourcesNotFindErr
		}
		log.GetLogger().Error("[repository] Merchant QueryMerChant FindOne", zap.Any("repo", repo.ToSql(ctx)), zap.Error(err))
		return nil, hcode.MgoExecErr
	}
	return
}

func (m *Merchant) QueryMerChants(ctx context.Context, repo repo.MerChantSpecificationRepo) (data []*entity.Merchant, err error) {
	return nil, nil
}
