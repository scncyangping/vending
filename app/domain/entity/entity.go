package entity

import "vending/app/types"

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
}

type BeneficiaryEn struct {
	Id     string                  `json:"id"`
	Type   string                  `json:"type"`   // 支付类型
	Status types.BeneficiaryStatus `json:"status"` // 状态：正常使用、停用、冻结
	Data   interface{}             `json:"data"`   // 支付使用数据：各个支付方式需要信息
	UserId string                  `json:"userId"` // 收款人Id,必是注册用户
}

type PaySubEn struct {
	Id        string          `json:"id" bson:"_id"`
	PayUser   string          `json:"payUser" bson:"payUser"`     // 支付人
	PayAmount float64         `json:"payAmount" bson:"payAmount"` // 支付金额
	BfId      string          `json:"bfId" bson:"bfId"`           // 支付关联ID
	PayStatus types.PayStatus `json:"payStatus" bson:"payStatus"` // 支付状态
	PayLog    []string        `json:"payLog" bson:"payLog"`       // 流转日志 ["已创建：支付url xxx","已支付，回调：xxx"]
}
