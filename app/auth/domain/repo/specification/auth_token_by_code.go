package specification

import (
	"context"
	"vending/app/auth/domain/repo"
	"vending/app/auth/infrastructure/pkg/hcode"
)

type AuthTokenByOpenId struct {
	OpenId string `json:"code"`
}

func NewAuthTokenSpecificationByoOenId(openId string) repo.AuthTokenSpecificationRepo {
	return &AuthTokenByOpenId{OpenId: openId}
}

func (m AuthTokenByOpenId) ParameterCheck(ctx context.Context) error {
	if m.OpenId == "" {
		return hcode.SysParameterErr
	}
	return nil
}

func (m AuthTokenByOpenId) ToSql(ctx context.Context) interface{} {
	return m.OpenId
}
