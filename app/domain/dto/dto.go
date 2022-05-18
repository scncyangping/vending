package dto

import (
	"vending/app/domain/obj"
	"vending/app/types"
)

type CommoditySaveReq struct {
	CommodityId  string  `json:"commodityId"`  //  商品分类Id
	Name         string  `json:"name"`         // 商品名称
	Amount       float64 `json:"amount"`       // 商品价格
	Des          string  `json:"des"`          // 描述
	Introduction string  `json:"introduction"` // 简介
	Type         uint8   `json:"type"`         // 商品类型
	ImageUrl     string  `json:"imageUrl"`     // 图片链接
	CategoryId   string  `json:"categoryId"`   //  商品分类Id
}

type CategorySaveReq struct {
	CategoryId string         `json:"categoryId"`
	Name       string         `json:"name"`
	PId        string         `json:"pId"`
	SellType   types.SellType `json:"sellType" ` // 0 一次性 1 可重复使用
}

type StockSaveReq struct {
	Data       interface{} `json:"data"`       // 库存内容
	CategoryId string      `json:"categoryId"` // 关联类别Id
}

type CreateTemItem struct {
	CommodityId string `json:"commodityId"` // 商品id
	Num         int    `json:"num"`         // 购买数量
}

// CreateTemOrder 下单参数
type CreateTemOrder struct {
	Commodities []CreateTemItem `json:"commodities"`
	PayBfId     string          `json:"pay_bf_id"`   // 支付方式id
	PayerSubObj obj.PayerSubObj `json:"payerSubObj"` // 下单额外信息
}

// TemOrderPayDto 临时订单支付信息
type TemOrderPayDto struct {
	Type types.BeneficiaryType `json:"type"`
	Data interface{}           `json:"data"` // 具体支付数据，比如支付宝当面付就为付款url地址
}
