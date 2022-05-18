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
}

// BeneficiaryEn 收款信息
type BeneficiaryEn struct {
	Id     string                  `json:"id"`
	Type   types.BeneficiaryType   `json:"type"`   // 支付类型
	Status types.BeneficiaryStatus `json:"status"` // 状态：正常使用、停用、冻结
	Data   interface{}             `json:"data"`   // 支付使用数据：各个支付方式需要信息
	UserId string                  `json:"userId"` // 收款人Id,必是注册用户
}

// PayDesEn 支付信息
type PayDesEn struct {
	Id          string          `json:"id"`
	PayUser     string          `json:"payUser"`     // 支付人
	PayAmount   float64         `json:"payAmount"`   // 支付金额
	PayStatus   types.PayStatus `json:"payStatus"`   // 支付状态
	PayLog      []string        `json:"payLog"`      // 流转日志 ["已创建：支付url xxx","已支付，回调：xxx"]
	PayerSubObj obj.PayerSubObj `json:"payerSubObj"` // 支付额外信息
}

// OrderEn 订单
type OrderEn struct {
	Id             string             `json:"id"`
	OriginalAmount float64            `json:"originalAmount"` // 总商品原金额
	Amount         float64            `json:"amount"`
	OrderStatus    types.OrderStatus  `json:"orderStatus"` // 订单状态 开始、待支付、完成
	Items          []obj.OrderItemObj `json:"items"`
}

// StockEn 库存
type StockEn struct {
	Id         string            `json:"id"`
	Data       interface{}       `json:"data"`       // 库存内容
	CategoryId string            `json:"categoryId"` // 关联类别Id
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
