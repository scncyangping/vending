package aggregate

import "vending/app/infrastructure/repository"

type Aggregate struct {
	CommodityAggregate *CommodityAggregate
	InventoryAggregate *InventoryAggregate
	OrderAggregate     *OrderAggregate
}

// NewAggregate wire
func NewAggregate(repo *repository.Repository) *Aggregate {
	return &Aggregate{
		InventoryAggregate: NewInventoryAggregate(repo.CategoryRepo, repo.StockRepo),
	}
}
