package impl

import (
	"vending/app/application/assembler"
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/dto"
	"vending/app/domain/aggregate/factory"
)

type CommoditySrvImp struct {
}

func NewCommoditySrvImp() *CommoditySrvImp {
	return &CommoditySrvImp{}
}

func (c *CommoditySrvImp) Save(cmd *cmd.CommoditySaveCmd) (string, error) {
	return factory.Instance.CommodityAggregate.SaveCommodity(cmd)
}

func (c *CommoditySrvImp) Update(cmd *cmd.CommodityUpdateCmd) error {
	return factory.Instance.CommodityAggregate.ModifyCommodity(cmd)
}

func (c *CommoditySrvImp) Delete(strings []string) error {
	return factory.Instance.CommodityAggregate.DeleteCommodityBatch(strings)
}

func (c *CommoditySrvImp) Get(s string) *dto.CommodityDto {
	if cg, err := factory.Instance.CommodityAggregateInstance(s); err != nil {
		return nil
	} else {
		return assembler.CommodityEnToDto(cg.CommodityEn)
	}
}

func (c *CommoditySrvImp) Up(commodityId string) error {
	if cg, err := factory.Instance.CommodityAggregateInstance(commodityId); err != nil {
		return nil
	} else {
		return cg.CommodityUp()
	}
}

func (c *CommoditySrvImp) Down(commodityId string) error {
	if cg, err := factory.Instance.CommodityAggregateInstance(commodityId); err != nil {
		return nil
	} else {
		return cg.CommodityDown()
	}
}

func (c *CommoditySrvImp) Query(query query.CommoditiesPageQuery) ([]*dto.CommodityDto, error) {
	return nil, nil
}
