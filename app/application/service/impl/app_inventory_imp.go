package impl

import (
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/dto"
	"vending/app/domain/aggregate/factory"
	"vending/app/domain/service"
	"vending/app/infrastructure/pkg/util"
	"vending/app/types"
	"vending/app/types/constants"
)

type InventorySrvImp struct {
	factory            *factory.AgFactory
	doInventoryService service.DoInventoryService
}

func NewInventorySrvImp(factory *factory.AgFactory, doInventoryService service.DoInventoryService) *InventorySrvImp {
	return &InventorySrvImp{factory: factory, doInventoryService: doInventoryService}
}

func (i *InventorySrvImp) SaveCategory(cmd *cmd.CategorySaveCmd) (string, error) {
	return i.doInventoryService.SaveCategory(cmd)
}

func (i *InventorySrvImp) UpdateCategory(cmd *cmd.CategoryUpdateCmd) error {
	if ig, err := i.factory.InventoryAggregateInstance(cmd.CategoryId); err != nil {
		return err
	} else {
		return ig.UpdateCategory(cmd)
	}
}

func (i *InventorySrvImp) RemoveCategoryByIds(strings []string) error {
	return i.doInventoryService.RemoveCategoryByIds(strings)
}

func (i *InventorySrvImp) InStockOne(cmd *cmd.StockSaveCmd) (string, error) {
	if ig, err := i.factory.InventoryAggregateInstance(cmd.CategoryId); err != nil {
		return constants.EmptyStr, err
	} else {
		return ig.InStockOne(cmd)
	}
}

func (i *InventorySrvImp) QueryCategory(query query.CategoryPageQuery) ([]*dto.CategoryDto, error) {
	qb := query.PageQuery.QBase()

	filter := types.B{}
	if query.Name != constants.EmptyStr {
		filter["name"] = types.B{"$reg": query.Name}
	}

	if dos, err := i.doInventoryService.QueryCategoryPageBy(qb.Skip, qb.Limit, qb.Sort, filter); err != nil {
		return nil, err
	} else {
		var dtoList []*dto.CategoryDto
		for _, v := range dos {
			dto := dto.CategoryDto{}
			util.StructCopy(dto, v)
			dtoList = append(dtoList, &dto)
		}
		return dtoList, nil
	}
}

func (i *InventorySrvImp) QueryStock(query query.StockPageQuery) ([]*dto.StockDto, error) {
	qb := query.PageQuery.QBase()

	filter := types.B{}
	if query.CategoryId != constants.EmptyStr {
		filter["categoryId"] = query.CategoryId
	}

	if dos, err := i.doInventoryService.QueryStockPageBy(qb.Skip, qb.Limit, qb.Sort, filter); err != nil {
		return nil, err
	} else {
		var dtoList []*dto.StockDto
		for _, v := range dos {
			dto := dto.StockDto{}
			util.StructCopy(dto, v)
			dtoList = append(dtoList, &dto)
		}
		return dtoList, nil
	}
}
