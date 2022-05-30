package commodity

import (
	"errors"
	"vending/app/application/cqe/cmd"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/util"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types/constants"
)

type DoCommoditySrvImpl struct {
	commodityRepo repo.CommodityRepo
	categoryRepo  repo.CategoryRepo
}

func NewDoCommoditySrvImpl(commodityRepo repo.CommodityRepo, categoryRepo repo.CategoryRepo) *DoCommoditySrvImpl {
	return &DoCommoditySrvImpl{commodityRepo: commodityRepo, categoryRepo: categoryRepo}
}

// QueryCommodityPageBy 分页查询
func (c *DoCommoditySrvImpl) QueryCommodityPageBy(skip, limit int64, sort, filter any) ([]*do.CommodityDo, error) {
	return c.commodityRepo.ListCommodityPageBy(skip, limit, sort, filter)
}

// QueryCommoditiesByIds 根据ids查询商品列表
func (c *DoCommoditySrvImpl) QueryCommoditiesByIds(ids []string) ([]*entity.CommodityEn, error) {
	var (
		comEns []*entity.CommodityEn
	)
	if comDos, err := c.commodityRepo.ListCommodityByIds(ids); err != nil {
		return nil, err
	} else {
		for _, v := range comDos {
			en := entity.CommodityEn{}
			util.StructCopy(v, &en)
			comEns = append(comEns, &en)
		}
	}
	return comEns, nil
}

// SaveCommodity 添加商品(包括商品基本信息及对应分类Id)
func (c *DoCommoditySrvImpl) SaveCommodity(req *cmd.CommoditySaveCmd) (string, error) {
	var (
		err         error
		en          entity.CommodityEn
		categoryDo  *do.CategoryDo
		commodityId = snowflake.NextId()
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
	en.Id = commodityId
	// 4. 保存
	if _, err = c.commodityRepo.SaveCommodity(&en, req.CategoryId); err != nil {
		return constants.EmptyStr, err
	}
	return en.Id, nil
}

// DeleteCommodityBatch 批量删除商品
func (c *DoCommoditySrvImpl) DeleteCommodityBatch(s []string) error {
	return c.commodityRepo.DeleteCommodityBatch(s)
}
