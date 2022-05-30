package factory

import (
	"vending/app/domain/aggregate"
	"vending/app/infrastructure/repository"
)

type AgFactory struct {
	commodityAggregate *aggregate.CommodityAggregate
	inventoryAggregate *aggregate.InventoryAggregate
	orderAggregate     *aggregate.OrderAggregate
	payAggregate       *aggregate.PayAggregate
}

// NewAggregate wire
func NewAggregate(repo *repository.Repository) *AgFactory {
	f := &AgFactory{
		commodityAggregate: aggregate.NewCommodityAggregate(repo.CommodityRepo, repo.CategoryRepo),
		inventoryAggregate: aggregate.NewInventoryAggregate(repo.CategoryRepo, repo.StockRepo),
		orderAggregate:     aggregate.NewOrderAggregate(repo.OrderRepo, repo.OrderTempRepo, repo.BeneficiaryRepo),
		payAggregate:       aggregate.NewPayAggregate(repo.BeneficiaryRepo),
	}
	return f
}

func (f *AgFactory) InventoryAggregateInstance(categoryId ...string) (*aggregate.InventoryAggregate, error) {
	return f.inventoryAggregate.Instance(categoryId...)
}

func (f *AgFactory) CommodityAggregateInstance(commodityId ...string) (*aggregate.CommodityAggregate, error) {
	return f.commodityAggregate.Instance(commodityId...)
}

func (f *AgFactory) OrderAggregateInstance(orderId ...string) (*aggregate.OrderAggregate, error) {
	return f.orderAggregate.Instance(orderId...)
}

func (f *AgFactory) PayAggregateInstance() (*aggregate.PayAggregate, error) {
	return f.payAggregate, nil
}
