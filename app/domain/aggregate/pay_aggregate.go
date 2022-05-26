package aggregate

import (
	"encoding/json"
	"vending/app/domain/entity"
	"vending/app/domain/obj"
	"vending/app/domain/repo"
	"vending/app/infrastructure/pkg/util/pay/alipay"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types"
	"vending/app/types/constants"
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

// Pay 生成支付信息
// 现仅对接支付宝面对面支付
// TODO 待扩展支付方式
func (o *PayAggregate) Pay(orderId string, amount float64) (string, error) {
	var (
		face, _ = alipay.NewAlipayFaceToFace("", "", "", "")
	)
	// 根据订单号及订单内容获取二维码
	orderAlipayFaceBody := obj.OrderAlipayFaceBody{}
	orderAlipayFaceBody.OutTradeNo = orderId
	orderAlipayFaceBody.Subject = "账号服务咨询" // TODO 订单标题
	orderAlipayFaceBody.TotalAmount = amount
	if qrcode, err := face.Start(orderAlipayFaceBody); err != nil {
		return constants.EmptyStr, err
	} else {
		return qrcode.(string), nil
	}
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
	en.Type = beneficiaryType
	en.Id = snowflake.NextId()
	return o.beneficiaryRepo.SaveBeneficiary(&en)
}
