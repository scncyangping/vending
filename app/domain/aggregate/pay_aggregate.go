package aggregate

import (
	"encoding/json"
	"vending/app/domain/dto"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/pkg/util/pay/alipay"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types"
	"vending/app/types/constants"
)

//

type payAggregateRepo struct {
	payDesRepo      repo.PayDesRepo
	beneficiaryRepo repo.BeneficiaryRepo
}

type PayAggregate struct {
	payAggregateRepo
	OrderId       string
	OrderEn       entity.OrderEn
	PayDesEn      entity.PayDesEn
	BeneficiaryEn entity.BeneficiaryEn
}

func NewPayAggregate(beneficiaryRepo repo.BeneficiaryRepo) *PayAggregate {
	return &PayAggregate{
		payAggregateRepo: payAggregateRepo{
			beneficiaryRepo: beneficiaryRepo,
		},
	}
}

// Pay 生成支付信息
func (o *PayAggregate) Pay(payRepo alipay.PayRepo) *dto.TemOrderPayDto {
	//payRepo.GetPayDetails()
	return nil
}

// SaveBeneficiary 添加收款方式,同一人同一类型支付方式可以绑定多个，绑定多个，收款时随机选择其中一个
func (o *PayAggregate) SaveBeneficiary(beneficiaryType types.BeneficiaryType, data any, userId string) (string, error) {
	var (
		en entity.BeneficiaryEn
	)
	// 转换为字符串存储
	if d, err := json.Marshal(data); err != nil {
		return constants.EmptyStr, err
	} else {
		en.Data = d
	}
	en.UserId = userId
	en.Type = beneficiaryType
	en.Id = snowflake.NextId()
	return o.beneficiaryRepo.SaveBeneficiary(&en)
}
