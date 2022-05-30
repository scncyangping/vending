package aggregate

import (
	"errors"
	"vending/app/application/cqe/cmd"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/pkg/log"
	"vending/app/infrastructure/pkg/util"
	"vending/app/types"
)

// 商品聚合，处理商品相关业务

type CommodityAggregateRepo struct {
	// 商品repo
	commodityRepo repo.CommodityRepo
	// 类别repo
	categoryRepo repo.CategoryRepo
}

type CommodityAggregate struct {
	CommodityAggregateRepo
	CommodityEn *entity.CommodityEn
	CommodityId string
}

func NewCommodityAggregate(commodityRepo repo.CommodityRepo, categoryRepo repo.CategoryRepo) *CommodityAggregate {
	return &CommodityAggregate{
		CommodityAggregateRepo: CommodityAggregateRepo{
			commodityRepo: commodityRepo,
			categoryRepo:  categoryRepo,
		},
	}
}

func (c *CommodityAggregate) Instance(CommodityId ...string) (*CommodityAggregate, error) {
	if len(CommodityId) < 1 {
		return c, nil
	}
	c.CommodityId = CommodityId[0]

	if ca, err := c.commodityRepo.GetCommodityById(c.CommodityId); err != nil {
		return c, err
	} else {
		util.StructCopy(&c.CommodityEn, ca)
	}
	return c, nil
}

// CommodityUp 商品上架
func (c *CommodityAggregate) CommodityUp() error {
	log.Logger().Infof("上架商品 [ %s ]", c.CommodityId)
	// 组装修改信息
	return c.commodityRepo.UpdateCommodity(types.B{"_id": c.CommodityId}, types.B{"status": types.CommodityUp})
}

// CommodityDown 商品下架
func (c *CommodityAggregate) CommodityDown() error {
	log.Logger().Infof("下架商品[ %s ] ", c.CommodityId)
	// TODO 商品下架 发送其他事件
	return c.commodityRepo.UpdateCommodity(types.B{"_id": c.CommodityId}, types.B{"status": types.CommodityDown})
}

// ModifyCommodity 修改商品基本信息
func (c *CommodityAggregate) ModifyCommodity(req *cmd.CommodityUpdateCmd) error {
	if m, err := util.StructToMap(req); err != nil {
		return errors.New("商品基本信息转换失败")
	} else {
		// 移除不需要的元素
		delete(m, "CommodityId")

		return c.commodityRepo.UpdateCommodity(types.B{"_id": c.CommodityId}, types.B{"$set": m})
	}
}

// DeleteCommodity 删除商品
func (c *CommodityAggregate) DeleteCommodity(CommodityId string) error {
	return c.commodityRepo.DeleteCommodity(CommodityId)
}
