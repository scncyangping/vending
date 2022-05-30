package cmd

import (
	"vending/app/domain/obj"
	"vending/app/types"
)

type LoginCmd struct {
	Name string `json:"name"` // 用户名
	Pwd  string `json:"pwd"`  // 密码
}

type RegisterCmd struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
}

type CommoditySaveCmd struct {
	Name         string  `json:"name"`         // 商品名称
	Amount       float64 `json:"amount"`       // 商品价格
	Des          string  `json:"des"`          // 描述
	Introduction string  `json:"introduction"` // 简介
	Type         uint8   `json:"type"`         // 商品类型
	ImageUrl     string  `json:"imageUrl"`     // 图片链接
	CategoryId   string  `json:"categoryId"`   // 商品分类Id
}

type CommodityUpdateCmd struct {
	CommoditySaveCmd
	CommodityId string `json:"commodityId"` // 商品Id
}

type CategorySaveCmd struct {
	Name     string         `json:"name"`      // 分类名称
	PId      string         `json:"pId"`       // 父分类Id
	SellType types.SellType `json:"sellType" ` // 库存使用类型 单次/重复
}

type CategoryUpdateCmd struct {
	CategorySaveCmd
	CategoryId string `json:"categoryId"` // 分类Id
}

type StockSaveCmd struct {
	Data       any    `json:"data"`       // 库存内容
	CategoryId string `json:"categoryId"` // 关联类别Id
}

//	CommodityId string `json:"commodityId"` // 商品Id
//	Num         int    `json:"num"`         // 购买数量
type CreateOrderCmd struct {
	Items        map[string]int        `json:"items"`   // 商品Id -> 数量
	PayDes       obj.PayDesObj         `json:"PayDes"`  // 额外信息
	PayType      types.BeneficiaryType `json:"payType"` // 支付类型
	CommodityIds []string              // 预构建商品ID
}

func (c *CreateOrderCmd) GetCommodityIds() []string {
	commodityIds := make([]string, 0)
	for key, _ := range c.Items {
		commodityIds = append(commodityIds, key)
	}
	return commodityIds
}
