package query

import "vending/app/types"

type PageQuery struct {
	Page   int              `json:"page"`   // 页
	Size   int              `json:"size"`   // 每页数量
	Sort   string           `json:"sort"`   // 排序字段
	SortBy types.SortByType `json:"sortBy"` // 正序/倒序
}

type PageQBase struct {
	Skip  int64
	Limit int64
	Sort  types.B
}

func (pq *PageQuery) QBase() *PageQBase {
	skip := 0
	if pq.Page > 0 {
		skip = (pq.Page - 1) * pq.Size
	}
	return &PageQBase{
		Skip:  int64(skip),
		Limit: int64(pq.Size),
		Sort:  types.B{pq.Sort: pq.SortBy},
	}
}

type OrderPageQuery struct {
	PageQuery
	CommodityName string `json:"commodityName"` // 商品名称
	CommodityId   string `json:"commodityId"`   // 商品Id
	PaymentPhone  string `json:"paymentPhone"`  // 支付人手机号
	PaymentEmail  string `json:"paymentEmail"`  // 支付人邮箱
}

type CommoditiesPageQuery struct {
	PageQuery
	Name       string `json:"name"`       // 商品名称
	CategoryId string `json:"categoryId"` // 分类Id
}

type CategoryPageQuery struct {
	PageQuery
	Name string `json:"name"` // 商品分类名称
}

type StockPageQuery struct {
	PageQuery
	CategoryId string `json:"categoryId"` // 关联类别Id
}
