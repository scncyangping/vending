package aggregate

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
)

type orderAggregateRepo struct {
	orderRepo     repo.OrderRepo
	orderTempRepo repo.OrderTempRepo
	payDesRepo    repo.PayDesRepo
}

type OrderAggregate struct {
	orderAggregateRepo
	OrderId       string
	OrderEn       entity.OrderEn
	PayDesEn      entity.PayDesEn
	BeneficiaryEn entity.BeneficiaryEn
}

func NewOrderAggregate(orderRepo repo.OrderRepo, orderTempRepo repo.OrderTempRepo,
	payDesRepo repo.PayDesRepo) *OrderAggregate {
	return &OrderAggregate{
		orderAggregateRepo: orderAggregateRepo{
			orderRepo:     orderRepo,
			orderTempRepo: orderTempRepo,
			payDesRepo:    payDesRepo,
		},
	}

}

// CreateTemOrder 创建临时订单
// 创建订单仅需支付信息金额
func (o *OrderAggregate) CreateTemOrder(commodityUserId string) {

}
