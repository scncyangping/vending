package factory

import (
	"vending/app/domain/aggregate"
	"vending/app/infrastructure/repository"
)

var Instance *Aggregate

type Aggregate struct {
	commodityAggregate *aggregate.CommodityAggregate
	inventoryAggregate *aggregate.InventoryAggregate
	orderAggregate     *aggregate.OrderAggregate
	payAggregate       *aggregate.PayAggregate
}

// NewAggregate wire
func NewAggregate(repo *repository.Repository) {
	f := &Aggregate{
		commodityAggregate: aggregate.NewCommodityAggregate(repo.CommodityRepo, repo.CategoryRepo),
		inventoryAggregate: aggregate.NewInventoryAggregate(repo.CategoryRepo, repo.StockRepo),
		orderAggregate:     aggregate.NewOrderAggregate(repo.OrderRepo, repo.OrderTempRepo, repo.BeneficiaryRepo),
		payAggregate:       aggregate.NewPayAggregate(repo.BeneficiaryRepo),
	}
	Instance = f
}

func (f *Aggregate) InventoryAggregateInstance(categoryId ...string) (*aggregate.InventoryAggregate, error) {
	return f.inventoryAggregate.Instance(categoryId...)
}

func (f *Aggregate) CommodityAggregateInstance(categoryId ...string) (*aggregate.CommodityAggregate, error) {
	return f.commodityAggregate.Instance(categoryId...)
}

func (f *Aggregate) OrderAggregateInstance(categoryId ...string) (*aggregate.OrderAggregate, error) {
	return f.orderAggregate.Instance(categoryId...)
}

func (f *Aggregate) PayAggregateInstance() (*aggregate.PayAggregate, error) {
	return f.payAggregate, nil
}
