package repo

import (
	"vending/app/domain/dto"
	"vending/app/types"
)

type PayRepo interface {
	GetPayDetails(types.BeneficiaryType, interface{}) *dto.TemOrderPayDto // 获取支付信息，比如：支付宝当面付，返回url地址
}
