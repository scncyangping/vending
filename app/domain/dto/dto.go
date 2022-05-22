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
	Data       any    `json:"data"`       // 库存内容
	CategoryId string `json:"categoryId"` // 关联类别Id
}

type CreateOrderReq struct {
	// 商品Id
	CommodityId string                `json:"commodityId"` // 商品Id
	Num         int                   `json:"num"`         // 购买数量
	PayDes      obj.PayDesObj         `json:"PayDes"`      // 额外信息
	PayType     types.BeneficiaryType `json:"payType"`
}
