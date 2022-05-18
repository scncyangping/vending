package vo

import "vending/app/types"

// StockVo 库存
type StockVo struct {
	Id         string            `json:"id"`
	Data       interface{}       `json:"data"`        // 库存内容
	CategoryId string            `json:"categoryId" ` // 关联类别Id
	OrderId    string            `json:"orderId"`     // 关联订单Id
	Status     types.StockStatus `json:"status"`      // 状态 0 待使用  1 已使用
}
