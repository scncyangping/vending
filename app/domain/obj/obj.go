package obj

// OrderItemObj 订单明细
type OrderItemObj struct {
	CommodityName  string  `json:"commodityName"`  // 商品类别+商品名称
	OriginalAmount float64 `json:"originalAmount"` // 商品原金额
	Amount         float64 `json:"amount"`         // 折扣计算后金额
}

// PayerSubObj 支付人信息
type PayerSubObj struct {
	Phone       string `json:"phone"`       // 电话号码
	Email       string `json:"email"`       // 邮箱
	PayerRemark string `json:"payerRemark"` // 支付备注
}
