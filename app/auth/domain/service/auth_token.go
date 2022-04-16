package service

import (
	"context"
	"vending/app/auth/domain/aggregate"
	"vending/app/auth/domain/dto"
	"vending/app/auth/domain/obj"
	"vending/app/auth/domain/repo"
	"vending/app/auth/domain/repo/specification"
	consts "vending/app/auth/infrastructure/conf"
	"vending/app/auth/infrastructure/pkg/hcode"
	"vending/app/auth/infrastructure/pkg/tool"
)

type AuthToken struct {
	factory       *aggregate.AuthFactory
	authTokenRepo repo.AuthTokenRepo
}

func (a *AuthToken) CreateToken(ctx context.Context, data dto.ProduceAuthTokenReq) (authTokenSimple dto.AuthTokenSimple, err error) {
	var (
		f aggregate.AuthTokenProduce
	)
	f, err = a.factory.NewProduceAuthToken(ctx, data)
	if err != nil {
		return
	}
	return f.ProduceToken(ctx)
}

func (a *AuthToken) RefreshAccessToken(ctx context.Context, data dto.RefreshAccessTokenReq) (authTokenSimple dto.AuthTokenSimple, err error) {
	if err = data.Check(); err != nil {
		return
	}
	var (
		token        = obj.AuthToken{}
		jwtToken     = tool.JwtToken{}
		jwtTokenData = tool.JwtTokenData{}
	)
	jwtTokenData, err = tool.CheckAuthToken(data.RefreshToken)
	if err != nil {
		return
	}
	token, err = a.authTokenRepo.QueryAuthToken(ctx, specification.NewAuthTokenSpecificationByoOenId(jwtTokenData.OpenId))
	if err != nil {
		return
	}
	jwtToken, err = tool.CreateAuthToken(jwtTokenData, consts.AuthAccessTokenCacheKeyTimeout)
	if err != nil {
		return
	}
	token.AccessToken = jwtToken.Token
	token.AccessTokenTimeline = jwtToken.TokenTimeline
	err = a.authTokenRepo.UpdateAuthToken(ctx, token)
	if err != nil {
		return
	}
	return token.TOSimple(), nil
}

func (a *AuthToken) GetUserInfo(ctx context.Context, data dto.OpenIdTokenReq) (user dto.UserSimple, err error) {
	if err = data.Check(); err != nil {
		return
	}
	var (
		f aggregate.AuthToken
	)
	f, err = a.factory.NewAuthToken(ctx, data)
	if err != nil {
		return
	}
	user, err = f.GetUserInfo(ctx)
	if err != nil {
		return
	}
	return
}

func (a *AuthToken) CheckToken(ctx context.Context, data dto.OpenIdTokenReq) (err error) {
	if err = data.Check(); err != nil {
		return
	}
	var (
		tokenData tool.JwtTokenData
	)
	if tokenData, err = tool.CheckAuthToken(data.Token); err != nil {
		return
	}
	if tokenData.OpenId != data.OpenId {
		return hcode.TokenValidErr
	}
	return
}
