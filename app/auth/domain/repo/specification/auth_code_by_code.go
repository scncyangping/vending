package specification

import (
	"context"
	"vending/app/auth/domain/repo"
	"vending/app/auth/infrastructure/pkg/hcode"
)

type AuthCodeByCode struct {
	Code string `json:"code"`
}

func NewAuthCodeSpecificationByCode(code string) repo.AuthCodeSpecificationRepo {
	return &AuthCodeByCode{Code: code}
}

func (m AuthCodeByCode) ParameterCheck(ctx context.Context) error {
	if m.Code == "" {
		return hcode.SysParameterErr
	}
	return nil
}

func (m AuthCodeByCode) ToSql(ctx context.Context) interface{} {
	return m.Code
}
