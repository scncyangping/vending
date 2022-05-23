package query

type PageQuery struct {
	Page   int    `json:"page"`   // 页
	Size   int    `json:"size"`   // 每页数量
	Sort   string `json:"sort"`   // 排序字段
	SortBy string `json:"sortBy"` // 正序/倒序
}

type OrderPageQuery struct {
	PageQuery
	CommodityName string `json:"commodityName"` // 商品名称
	PaymentPhone  string `json:"paymentPhone"`  // 支付人手机号
	PaymentEmail  string `json:"paymentEmail"`  // 支付人邮箱
}

type CommoditiesPageQuery struct {
	PageQuery
	Name string `json:"name"` // 商品名称
}

type CategoryPageQuery struct {
	PageQuery
	Name string `json:"name"` // 商品分类名称
}

type StockPageQuery struct {
	PageQuery
	CategoryId string `json:"categoryId"` // 关联类别Id
}
