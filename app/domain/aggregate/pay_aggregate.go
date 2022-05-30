package aggregate

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
)

// 创建支付
// 获取支付状态
// 添加支付方式
// 取消支付

type payAggregateRepo struct {
	beneficiaryRepo repo.BeneficiaryRepo
	// paymentRepo     pay.PaymentRepo 暂不对接其他支付方式
}

type PayAggregate struct {
	payAggregateRepo
	BeneficiaryEn entity.BeneficiaryEn
}

func NewPayAggregate(beneficiaryRepo repo.BeneficiaryRepo) *PayAggregate {
	return &PayAggregate{
		payAggregateRepo: payAggregateRepo{
			beneficiaryRepo: beneficiaryRepo,
		},
	}
}
