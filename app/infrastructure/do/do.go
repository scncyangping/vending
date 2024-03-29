package do

import (
	"vending/app/types"
)

type Do struct {
	Id         string `json:"id" bson:"_id"`
	CreateTime int64  `json:"createTime" bson:"createTime"`
	UpdateTime int64  `json:"updateTime" bson:"updateTime"`
	IsDeleted  uint8  `json:"isDeleted" bson:"isDeleted"`
	Remark     string `json:"remark" bson:"remark"`
}

// UserDo 用户
type UserDo struct {
	Do
	Name     string           `json:"name" bson:"name"`
	NickName string           `json:"nickName" bson:"nickName"`
	Phone    string           `json:"phone" bson:"phone"`
	Email    string           `json:"email" bson:"email"`
	Pwd      string           `json:"pwd" bson:"pwd"`
	Roles    []RoleDo         `json:"roles" bson:"roles"`
	Status   types.UserStatus `json:"status"bson:"status"`
}

// RoleDo 角色
type RoleDo struct {
	Do
	Name string `json:"name" bson:"name"` // 角色名称
}

// CommodityDo 商品
type CommodityDo struct {
	Do
	Name         string                `json:"name" bson:"name"`                 // 商品名称
	Amount       float64               `json:"amount" bson:"amount"`             // 商品价格
	Des          string                `json:"des" bson:"name"`                  // 描述
	Introduction string                `json:"introduction" bson:"introduction"` // 简介
	Type         uint8                 `json:"type" bson:"type"`                 // 商品类型
	ImageUrl     string                `json:"imageUrl" bson:"imageUrl"`         // 图片链接
	Status       types.CommodityStatus `json:"status" bson:"status"`             // 商品状态
	CategoryId   string                `json:"categoryId" bson:"categoryId"`     // 类别Id
	OwnerId      string                `json:"ownerId" bson:"ownerId"`           // 商品拥有人,可转移,转移后收款方式改为转移人
}

type BeneficiarySub struct {
	Type    types.BeneficiaryType   `json:"type" bson:"type"`       // 支付类型
	Status  types.BeneficiaryStatus `json:"status" bson:"status"`   // 状态：正常使用、停用、冻结
	Data    any                     `json:"data" bson:"data"`       // 支付使用数据：各个支付方式需要信息
	OwnerId string                  `json:"ownerId" bson:"ownerId"` // 收款人Id
}

// BeneficiaryDo 收款信息
type BeneficiaryDo struct {
	Do
	BeneficiarySub
}

// OrderDo 订单
type OrderDo struct {
	Do
	OriginalAmount float64           `json:"originalAmount" bson:"originalAmount"` // 总商品原金额
	Amount         float64           `json:"amount" bson:"amount"`                 // 总商品折扣金额
	Items          []OrderItemSub    `json:"items" bson:"items"`                   // 订单明细
	Payment        PayDesSubDo       `json:"payment" bson:"payment"`               // 支付信息
	OrderStatus    types.OrderStatus `json:"orderStatus" bson:"orderStatus"`       // 订单状态 开始、待支付、完成
	BfDo           BeneficiarySub    `json:"bf" bson:"bf"`                         // 支付关联信息
}

// StockDo 库存
type StockDo struct {
	Do
	Data       any               `json:"data" bson:"data"`             // 库存内容
	CategoryId string            `json:"categoryId" bson:"categoryId"` // 关联类别Id
	OrderId    string            `json:"orderId" bson:"orderId"`       // 关联订单Id
	Status     types.StockStatus `json:"status" bson:"status"`         // 状态 1 待使用  2 已使用
}

// CategoryDo 类别
type CategoryDo struct {
	Do
	Name     string         `json:"name" bson:"name"`         // 类别名称
	PId      string         `json:"pId" bson:"pId"`           // 父类别Id
	StockNum int            `json:"stockNum" bson:"stockNum"` // 库存数量 用库存数量去锁定待支付订单
	SellType types.SellType `json:"sellType" bson:"sellType"` // 1 一次性 2 可重复使用
}

// OrderItemSub 订单明细
type OrderItemSub struct {
	CommodityId    string         `json:"commodityId" bson:"commodityId"`       // 商品Id
	CommodityName  string         `json:"commodityName" bson:"commodityName"`   // 商品类别+商品名称
	OriginalAmount float64        `json:"originalAmount" bson:"originalAmount"` // 商品原金额
	Amount         float64        `json:"amount" bson:"amount"`                 // 折扣计算后金额
	OwnerId        string         `json:"ownerId"`                              // 商品拥有人,可转移,转移后收款方式改为转移人
	Payment        BeneficiarySub `json:"payment" bson:"amount"`                // 商品关联收款信息
}

// PayDesSubDo 支付信息 作为订单的元素，不单独存储
type PayDesSubDo struct {
	Phone       string          `json:"phone" bson:"phone"`             // 电话号码
	Email       string          `json:"email" bson:"email"`             // 邮箱
	PayerRemark string          `json:"payerRemark" bson:"payerRemark"` // 支付备注
	PayStatus   types.PayStatus `json:"payStatus" bson:"payStatus"`     // 支付状态
	PayLog      []string        `json:"payLog" bson:"payLog"`           // 流转日志 ["已创建：支付url xxx","已支付，回调：xxx"]
}
