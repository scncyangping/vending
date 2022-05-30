package impl

import (
	"vending/app/application/assembler"
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/dto"
	"vending/app/domain/aggregate/factory"
	"vending/app/domain/service"
	"vending/app/infrastructure/pkg/util"
	"vending/app/types"
	"vending/app/types/constants"
)

type CommoditySrvImp struct {
	factory            *factory.AgFactory
	doCommodityService service.DoCommodityService
}

func NewCommoditySrvImp(factory *factory.AgFactory, doCommodityService service.DoCommodityService) *CommoditySrvImp {
	return &CommoditySrvImp{factory: factory, doCommodityService: doCommodityService}
}

func (c *CommoditySrvImp) Save(cmd *cmd.CommoditySaveCmd) (string, error) {
	return c.doCommodityService.SaveCommodity(cmd)
}

func (c *CommoditySrvImp) Update(cmd *cmd.CommodityUpdateCmd) error {
	if cg, err := c.factory.CommodityAggregateInstance(cmd.CommodityId); err != nil {
		return err
	} else {
		return cg.ModifyCommodity(cmd)
	}
}

func (c *CommoditySrvImp) Delete(strings []string) error {
	return c.doCommodityService.DeleteCommodityBatch(strings)
}

func (c *CommoditySrvImp) Get(s string) *dto.CommodityDto {
	if cg, err := c.factory.CommodityAggregateInstance(s); err != nil {
		return nil
	} else {
		return assembler.CommodityEnToDto(cg.CommodityEn)
	}
}

func (c *CommoditySrvImp) Up(commodityId string) error {
	if cg, err := c.factory.CommodityAggregateInstance(commodityId); err != nil {
		return nil
	} else {
		return cg.CommodityUp()
	}
}

func (c *CommoditySrvImp) Down(commodityId string) error {
	if cg, err := c.factory.CommodityAggregateInstance(commodityId); err != nil {
		return nil
	} else {
		return cg.CommodityDown()
	}
}

func (c *CommoditySrvImp) Query(query query.CommoditiesPageQuery) ([]*dto.CommodityDto, error) {
	qb := query.PageQuery.QBase()

	filter := types.B{}
	if query.Name != constants.EmptyStr {
		filter["name"] = types.B{"$reg": query.Name}
	}
	if query.CategoryId != constants.EmptyStr {
		filter["categoryId"] = types.B{"$reg": query.CategoryId}
	}

	if dos, err := c.doCommodityService.QueryCommodityPageBy(qb.Skip, qb.Limit, qb.Sort, filter); err != nil {
		return nil, err
	} else {
		var dtoList []*dto.CommodityDto
		for _, v := range dos {
			dto := dto.CommodityDto{}
			util.StructCopy(dto, v)
			dtoList = append(dtoList, &dto)
		}
		return dtoList, nil
	}
}
