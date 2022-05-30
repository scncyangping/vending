package pay

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

type DoPaySrvImpl struct {
	beneficiaryRepo repo.BeneficiaryRepo
}

func NewDoPaySrvImpl(beneficiaryRepo repo.BeneficiaryRepo) *DoPaySrvImpl {
	return &DoPaySrvImpl{beneficiaryRepo: beneficiaryRepo}
}

// PayUrl 生成支付信息
// 现仅对接支付宝面对面支付
// TODO 待扩展支付方式
func (o *DoPaySrvImpl) PayUrl(orderId string, amount float64) (string, error) {
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
func (o *DoPaySrvImpl) SaveBeneficiary(beneficiaryType types.BeneficiaryType, data any, userId string) (string, error) {
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
	en.OwnerId = userId
	return o.beneficiaryRepo.SaveBeneficiary(&en)
}
