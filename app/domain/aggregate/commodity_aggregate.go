package aggregate

import "vending/app/domain/entity"

// 商品聚合 包含商品关联行为

// 导入商品 关联库存
// 修改商品信息
//

type CommodityAggregateRepo interface {
}

type CommodityAggregate struct {
	Cma        entity.CommodityEn
	CategoryId string
	repo       CommodityAggregateRepo
}
