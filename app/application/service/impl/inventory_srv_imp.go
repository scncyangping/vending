package impl

import (
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/dto"
	"vending/app/domain/aggregate/factory"
)

type InventorySrvImp struct {
}

func NewInventorySrvImp() *InventorySrvImp {
	return &InventorySrvImp{}
}

func (i *InventorySrvImp) SaveCategory(cmd *cmd.CategorySaveCmd) (string, error) {
	return factory.Instance.InventoryAggregate.SaveCategory(cmd)
}

func (i *InventorySrvImp) UpdateCategory(cmd *cmd.CategoryUpdateCmd) error {
	return factory.Instance.InventoryAggregate.UpdateCategory(cmd)
}

func (i *InventorySrvImp) RemoveCategoryByIds(strings []string) error {
	return factory.Instance.InventoryAggregate.RemoveCategoryByIds(strings)
}

func (i *InventorySrvImp) InStockOne(cmd *cmd.StockSaveCmd) (string, error) {
	return factory.Instance.InventoryAggregate.InStockOne(cmd)
}

func (i *InventorySrvImp) QueryCategory(query query.CategoryPageQuery) ([]*dto.CategoryDto, error) {
	return nil, nil
}

func (i *InventorySrvImp) QueryStock(query query.StockPageQuery) ([]*dto.StockDto, error) {
	return nil, nil
}
