package dto

import (
	"vending/app/domain/obj"
	"vending/app/types"
)

type BaseDto struct {
	Id         string `json:"id"`
	CreateTime int64  `json:"createTime"`
}

type CommodityDto struct {
	BaseDto
	Name         string  `json:"name"`         // 商品名称
	Amount       float64 `json:"amount"`       // 商品价格
	Des          string  `json:"des"`          // 描述
	Introduction string  `json:"introduction"` // 简介
	Type         uint8   `json:"type"`         // 商品类型
	ImageUrl     string  `json:"imageUrl"`     // 图片链接
	CategoryId   string  `json:"categoryId"`   // 商品分类Id
}

type CategoryDto struct {
	BaseDto
	Name     string         `json:"name"`
	PId      string         `json:"pId"`
	SellType types.SellType `json:"sellType" `
}

type StockDto struct {
	BaseDto
	Data       any    `json:"data"`       // 库存内容
	CategoryId string `json:"categoryId"` // 关联类别Id
}

type OrderDto struct {
	BaseDto
	OriginalAmount float64            `json:"originalAmount"` // 总商品原金额
	Amount         float64            `json:"amount"`         // 总商品折扣金额
	Items          []obj.OrderItemObj `json:"items"`          // 订单明细
	Payment        obj.PayDesObj      `json:"payment"`        // 支付信息
	OrderStatus    types.OrderStatus  `json:"orderStatus"`    // 订单状态 开始、待支付、完成
	BfDto          obj.BeneficiaryObj `json:"bf"`             // 支付关联信息
}

type OrderListDto struct {
	BaseDto
	Amount      float64           `json:"amount"`      // 总商品折扣金额
	OrderStatus types.OrderStatus `json:"orderStatus"` // 订单状态 开始、待支付、完成
}

type LoginDto struct {
	Token string `json:"token"`
}

type UserDto struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Type       uint8  `json:"type"`
	Status     uint8  `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Token      string `json:"token"`
}
