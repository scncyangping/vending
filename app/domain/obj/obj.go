package obj

import "vending/app/types"

// OrderItemObj 订单明细
type OrderItemObj struct {
	CommodityId    string         `json:"commodityId"`           // 商品Id 不可变数据
	CommodityName  string         `json:"commodityName"`         // 商品类别+商品名称
	OriginalAmount float64        `json:"originalAmount"`        // 商品原金额
	Amount         float64        `json:"amount"`                // 折扣计算后金额
	OwnerId        string         `json:"ownerId"`               // 商品拥有人,可转移,转移后收款方式改为转移人
	Payment        BeneficiaryObj `json:"payment" bson:"amount"` // 商品关联收款信息
}

type BeneficiaryObj struct {
	Type    types.BeneficiaryType   `json:"type"`    // 支付类型
	Status  types.BeneficiaryStatus `json:"status"`  // 状态：正常使用、停用、冻结
	Data    any                     `json:"data"`    // 支付使用数据：各个支付方式需要信息
	OwnerId string                  `json:"ownerId"` // 收款人Id
}

// PayDesObj 支付信息
type PayDesObj struct {
	Phone       string          `json:"phone"`       // 电话号码
	Email       string          `json:"email"`       // 邮箱
	PayerRemark string          `json:"payerRemark"` // 支付备注
	PayStatus   types.PayStatus `json:"payStatus" `  // 支付状态
	PayLog      []string        `json:"payLog" `     // 流转日志 ["已创建：支付url xxx","已支付，回调：xxx"]
}

// OrderAlipayFaceBody 支付宝面对面支付
type OrderAlipayFaceBody struct {
	OutTradeNo  string  `json:"out_trade_no"` // 商户订单号
	Subject     string  `json:"subject"`      // 订单标题
	TotalAmount float64 `json:"total_amount"` // 总金额
}
