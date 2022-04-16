package specification

import (
	"context"
	"vending/app/auth/domain/repo"
	"vending/app/auth/infrastructure/pkg/hcode"
)

type MerchantByAppid struct {
	APPID string `json:"appid" bson:"appid"`
}

func NewMerchantSpecificationByAPPID(APPID string) repo.MerChantSpecificationRepo {
	return &MerchantByAppid{APPID: APPID}
}

func (m MerchantByAppid) ParameterCheck(ctx context.Context) error {
	if m.APPID == "" {
		return hcode.SysParameterErr
	}
	return nil
}

func (m MerchantByAppid) ToSql(ctx context.Context) interface{} {
	return map[string]string{
		"appid": m.APPID,
	}
}
