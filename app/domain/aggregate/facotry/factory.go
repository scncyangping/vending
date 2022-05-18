package facotry

import "vending/app/domain/aggregate"

type Factory struct {
	aggregate *aggregate.Aggregate
}

func NewFactory(agg *aggregate.Aggregate) *Factory {
	return &Factory{
		aggregate: agg,
	}
}
func (f *Factory) InventoryAggregate(categoryId ...string) (*aggregate.InventoryAggregate, error) {
	return f.aggregate.InventoryAggregate.Instance(categoryId...)
}
