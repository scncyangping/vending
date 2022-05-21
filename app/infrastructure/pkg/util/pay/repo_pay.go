package pay

import "vending/app/types"

type PaymentRepo interface {
	Start(any) (any, error)                 // 获取支付信息，比如：支付宝当面付，返回url地址
	Status(string) (types.PayStatus, error) // 商户订单号主动查询支付状态
	Notify(any) (bool, error)               // 支付回调
}
