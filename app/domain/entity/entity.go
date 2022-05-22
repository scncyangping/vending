package entity

import (
	"vending/app/domain/obj"
	"vending/app/types"
)

type RoleEn struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Code uint8  `json:"code"`
}

type UserEn struct {
	Id       string           `json:"id"`
	Name     string           `json:"name"`
	NickName string           `json:"nickName"`
	Phone    string           `json:"phone"`
	Email    string           `json:"email"`
	Pwd      string           `json:"pwd"`
	Status   types.UserStatus `json:"status"`
}

type CommodityEn struct {
	Id           string                `json:"id"`
	Name         string                `json:"name"`         // 商品名称
	Amount       float64               `json:"amount"`       // 商品价格
	Des          string                `json:"des"`          // 描述
	Introduction string                `json:"introduction"` // 简介
	Type         uint8                 `json:"type"`         // 商品类型
	ImageUrl     string                `json:"imageUrl"`     // 图片链接
	Status       types.CommodityStatus `json:"status"`       // 商品状态
	CategoryId   string                `json:"categoryId"`   //  商品分类Id
	OwnerId      string                `json:"ownerId"`      // 商品拥有人,可转移,转移后收款方式改为转移人
}

// BeneficiaryEn 收款信息
type BeneficiaryEn struct {
	Id string `json:"id"`
	obj.BeneficiaryObj
}

// StockEn 库存
type StockEn struct {
	Id         string            `json:"id"`
	Data       any               `json:"data"`       // 库存内容
	CategoryId string            `json:"categoryId"` // 关联类别Id
	OrderId    string            `json:"orderId"`    // 关联订单Id
	Status     types.StockStatus `json:"status"`
}

// CategoryEn 类别
type CategoryEn struct {
	Id       string         `json:"id"`
	Name     string         `json:"name" `     // 类别名称
	PId      string         `json:"pId" `      // 父类别Id
	StockNum int            `json:"stockNum" ` // 库存数量 用库存数量去锁定待支付订单
	SellType types.SellType `json:"sellType" ` // 0 一次性 1 可重复使用
}

// OrderEn 订单
type OrderEn struct {
	Id             string             `json:"id"`
	OriginalAmount float64            `json:"originalAmount" bson:"originalAmount"` // 总商品原金额
	Amount         float64            `json:"amount" bson:"amount"`                 // 总商品折扣金额
	Items          []obj.OrderItemObj `json:"items" bson:"items"`                   // 订单明细
	PayDesObj      obj.PayDesObj      `json:"payment" bson:"payment"`
	OrderStatus    types.OrderStatus  `json:"orderStatus" bson:"orderStatus"` // 订单状态 开始、待支付、完成
	BfObj          obj.BeneficiaryObj `json:"bf" bson:"bf"`
}
