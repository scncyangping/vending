package obj

// OrderItemObj 订单明细
type OrderItemObj struct {
	CommodityName  string  `json:"commodityName"`  // 商品类别+商品名称
	OriginalAmount float64 `json:"originalAmount"` // 商品原金额
	Amount         float64 `json:"amount"`         // 折扣计算后金额
}
