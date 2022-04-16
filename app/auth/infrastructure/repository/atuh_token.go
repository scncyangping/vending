package repository

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"vending/app/auth/domain/obj"
	"vending/app/auth/domain/repo"
	consts "vending/app/auth/infrastructure/conf"
	"vending/app/auth/infrastructure/pkg/hcode"
	"vending/app/auth/infrastructure/pkg/log"
)

var _ repo.AuthTokenRepo = (*AuthToken)(nil)

type AuthToken struct {
	repository
}

func (a *AuthToken) getCacheKey(data string) string {
	return fmt.Sprintf("%s%s", consts.AuthTokenCacheKey, data)
}

func (a *AuthToken) CreateAuthToken(ctx context.Context, data obj.AuthToken) error {
	saveData, err := Marshal(&data)
	if err != nil {
		log.GetLogger().Error("[AuthToken] CreateAuthToken Marshal", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	if err := a.rds.Set(a.getCacheKey(data.OpenID), saveData, consts.AuthRefreshTokenCacheKeyTimeout).Err(); err != nil {
		log.GetLogger().Error("[AuthToken] CreateAuthToken Set", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	return nil
}

func (a *AuthToken) UpdateAuthToken(ctx context.Context, data obj.AuthToken) error {
	saveData, err := Marshal(&data)
	if err != nil {
		log.GetLogger().Error("[AuthToken] UpdateAuthToken Marshal", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	if err := a.rds.Set(a.getCacheKey(data.OpenID), saveData, consts.AuthRefreshTokenCacheKeyTimeout).Err(); err != nil {
		log.GetLogger().Error("[AuthToken] UpdateAuthToken Set", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	return nil
}

func (a *AuthToken) QueryAuthToken(ctx context.Context, repo repo.AuthTokenSpecificationRepo) (obj.AuthToken, error) {
	data, err := a.rds.Get(a.getCacheKey(fmt.Sprint(repo.ToSql(ctx)))).Bytes()
	if err != nil {
		log.GetLogger().Error("[QueryAuthToken] Get", zap.Any("req", repo.ToSql(ctx)), zap.Error(err))
		return obj.AuthToken{}, hcode.RedisExecErr
	}
	var authToken obj.AuthToken
	if err := Unmarshal(data, &authToken); err != nil {
		log.GetLogger().Error("[QueryAuthToken] Unmarshal", zap.Any("req", repo.ToSql(ctx)), zap.Error(err))
		return obj.AuthToken{}, hcode.RedisExecErr
	}
	return authToken, nil
}
