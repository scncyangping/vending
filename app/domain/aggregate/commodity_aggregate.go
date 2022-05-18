package aggregate

import (
	"errors"
	"vending/app/domain/dto"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/log"
	"vending/app/infrastructure/pkg/util"
	"vending/app/types"
	"vending/app/types/constants"
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
	commodityEn *entity.CommodityEn
	commodityId string
}

func NewCommodityAggregate(commodityRepo repo.CommodityRepo, categoryRepo repo.CategoryRepo) *CommodityAggregate {
	return &CommodityAggregate{
		CommodityAggregateRepo: CommodityAggregateRepo{
			commodityRepo: commodityRepo,
			categoryRepo:  categoryRepo,
		},
	}
}

func (c *CommodityAggregate) Instance(commodityId ...string) (*CommodityAggregate, error) {
	if len(commodityId) < 1 {
		return c, nil
	}
	c.commodityId = commodityId[0]

	if ca, err := c.commodityRepo.GetCommodityById(c.commodityId); err != nil {
		return c, err
	} else {
		util.StructCopy(&c.commodityEn, ca)
	}
	return c, nil
}

// CommodityUp 商品上架
func (c *CommodityAggregate) CommodityUp() error {
	log.Logger().Infof("上架商品 [ %s ]", c.commodityId)
	// 组装修改信息
	return c.commodityRepo.UpdateCommodity(types.B{"_id": c.commodityId}, types.B{"status": types.CommodityUp})
}

// CommodityDown 商品下架
func (c *CommodityAggregate) CommodityDown(status types.CommodityStatus) error {
	log.Logger().Infof("下架商品[ %s ] ", c.commodityId)
	// TODO 商品下架 发送其他事件
	return c.commodityRepo.UpdateCommodity(types.B{"_id": c.commodityId}, types.B{"status": types.CommodityDown})
}

// ModifyCommodity 修改商品基本信息
func (c *CommodityAggregate) ModifyCommodity(req *dto.CommoditySaveReq) error {
	if m, err := util.StructToMap(req); err != nil {
		return errors.New("商品基本信息转换失败")
	} else {
		return c.commodityRepo.UpdateCommodity(types.B{"_id": c.commodityId}, types.B{"$set": m})
	}
}

// SaveCommodity 添加商品(包括商品基本信息及对应分类Id)
func (c *CommodityAggregate) SaveCommodity(req *dto.CommoditySaveReq) (string, error) {
	var (
		err        error
		en         entity.CommodityEn
		categoryDo *do.CategoryDo
	)
	// 1. 查询对应分类是否存在
	if categoryDo, err = c.categoryRepo.GetCategoryById(req.CategoryId); err != nil {
		return constants.EmptyStr, err
	} else {
		if categoryDo == nil {
			return constants.EmptyStr, errors.New("对应商品分类不存在")
		}
	}
	// 2. 组装基本信息
	util.StructCopy(&en, categoryDo)
	// 3. 添加唯一Id
	en.Id = c.commodityId
	// 4. 保存
	if _, err = c.commodityRepo.SaveCommodity(&en, req.CategoryId); err != nil {
		return constants.EmptyStr, err
	}
	c.commodityEn = &en
	return c.commodityId, nil
}
