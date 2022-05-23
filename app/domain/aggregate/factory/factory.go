package factory

import (
	"vending/app/domain/aggregate"
	"vending/app/infrastructure/repository"
)

var Instance *factoryAggregate

type factoryAggregate struct {
	CommodityAggregate *aggregate.CommodityAggregate
	InventoryAggregate *aggregate.InventoryAggregate
	OrderAggregate     *aggregate.OrderAggregate
	PayAggregate       *aggregate.PayAggregate
}

// NewAggregate wire
func NewAggregate(repo *repository.Repository) {
	f := &factoryAggregate{
		CommodityAggregate: aggregate.NewCommodityAggregate(repo.CommodityRepo, repo.CategoryRepo),
		InventoryAggregate: aggregate.NewInventoryAggregate(repo.CategoryRepo, repo.StockRepo),
		OrderAggregate:     aggregate.NewOrderAggregate(repo.OrderRepo, repo.OrderTempRepo, repo.BeneficiaryRepo),
		PayAggregate:       aggregate.NewPayAggregate(repo.BeneficiaryRepo),
	}
	Instance = f
}

func (f *factoryAggregate) InventoryAggregateInstance(categoryId ...string) (*aggregate.InventoryAggregate, error) {
	return f.InventoryAggregate.Instance(categoryId...)
}

func (f *factoryAggregate) CommodityAggregateInstance(categoryId ...string) (*aggregate.CommodityAggregate, error) {
	return f.CommodityAggregate.Instance(categoryId...)
}

func (f *factoryAggregate) OrderAggregateInstance(categoryId ...string) (*aggregate.OrderAggregate, error) {
	return f.OrderAggregate.Instance(categoryId...)
}

func (f *factoryAggregate) PayAggregateInstance() (*aggregate.PayAggregate, error) {
	return f.PayAggregate, nil
}
