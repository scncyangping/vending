package aggregate

import (
	"errors"
	"vending/app/domain/dto"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/domain/vo"
	"vending/app/infrastructure/pkg/util"
	"vending/app/types"
	"vending/app/types/constants"
)

type InventoryAggregateRepo struct {
	// 类别repo
	categoryRepo repo.CategoryRepo
	// 库存repo
	stockRepo repo.StockRepo
}

type InventoryAggregate struct {
	InventoryAggregateRepo
	categoryEn entity.CategoryEn
	stockEn    []entity.StockEn
	categoryId string
}

func NewInventoryAggregate(categoryRepo repo.CategoryRepo, stockRepo repo.StockRepo) *InventoryAggregate {
	return &InventoryAggregate{
		InventoryAggregateRepo: InventoryAggregateRepo{
			categoryRepo: categoryRepo,
			stockRepo:    stockRepo,
		},
	}
}

func (c *InventoryAggregate) Instance(categoryId ...string) (*InventoryAggregate, error) {
	if len(categoryId) < 1 {
		return c, nil
	}
	c.categoryId = categoryId[0]

	if ca, err := c.categoryRepo.GetCategoryById(c.categoryId); err != nil {
		return c, err
	} else {
		util.StructCopy(&c.categoryEn, ca)
	}
	return c, nil
}

// OutStock 出库
func (c *InventoryAggregate) OutStock(num int) ([]*vo.StockVo, error) {
	var (
		stockVos []*vo.StockVo
		stockIds []string
	)
	// 校验库存
	if c.categoryEn.StockNum < num {
		return nil, errors.New("库存不足")
	} else {
		c.categoryEn.StockNum -= num
	}
	// 取出指定数量库存
	if stocks, err := c.stockRepo.ListStockPageBy(0, int64(num), types.B{"createTime": -1}, types.B{"status": types.StockNormal}); err != nil {
		return nil, err
	} else {
		// 整理数据
		if len(stocks) != num {
			return nil, errors.New("库存不足")
		}
		for _, v := range stocks {
			stockIds = append(stockIds, v.Id)
			vo := vo.StockVo{}
			util.StructCopy(vo, v)
			stockVos = append(stockVos, &vo)
		}
	}
	// 更新取出数据状态为已使用
	if count, err := c.stockRepo.UpdateStock(types.B{"_id": types.B{"$in": stockIds}}, types.B{"$set": types.B{"status": types.StockNormal}}); err != nil {
		return nil, errors.New("扣减库存失败")
	} else {
		if int(count) != num {
			// 说明在修改库存的时候被其他协程修改了
			// 重新计算
			return c.OutStock(num)
		}
	}
	// 修改分类统计库存总数量
	err := c.stockNum()
	if err != nil {
		return stockVos, errors.New("出库失败")
	}
	return stockVos, nil
}

// InStockOne 入库
func (c *InventoryAggregate) InStockOne(dto *dto.StockSaveReq) (string, error) {
	var (
		en entity.StockEn
	)
	// 校验分类是否存在
	if !c.existCategoryById(dto.CategoryId) {
		return constants.EmptyStr, errors.New("对应分类不存在")
	}
	en.CategoryId = dto.CategoryId
	en.Status = types.StockNormal
	en.Data = dto.Data

	// 修改分类统计库存总数量
	err := c.stockNum()
	if err != nil {
		return constants.EmptyStr, errors.New("入库失败")
	}
	return c.stockRepo.SaveStock(&en)
}

// RemoveCategoryByIds 批量移除分类
func (c *InventoryAggregate) RemoveCategoryByIds(ids []string) error {
	// 查询分类下是否有库存数据，若有，不允许移除
	if l, err := c.stockRepo.ListStockBy(
		types.B{"categoryId": types.B{"$in": ids}, "Status": types.StockNormal}); err != nil {
		return err
	} else {
		if len(l) > 0 {
			return errors.New("该分类下存在未使用库存,不允许删除")
		}
	}
	return c.categoryRepo.DeleteCategoryByIds(ids)
}

// UpdateCategory 修改分类基本信息
func (c *InventoryAggregate) UpdateCategory(req *dto.CategorySaveReq) error {
	// 查询对应分类是否存在
	if c.existCategoryByName(req.Name) {
		return errors.New("该分类已存在")
	}

	m := map[string]any{}

	if req.Name != constants.EmptyStr {
		m["name"] = req.Name
	}
	if req.PId != constants.EmptyStr {
		m["pId"] = req.PId
	}
	if req.SellType != 0 {
		m["sellType"] = req.SellType
	}

	return c.categoryRepo.UpdateCategory(types.B{"_id": req.CategoryId}, types.B{"$set": m})
}

// stockNum 分类统计库存数量修改
func (c *InventoryAggregate) stockNum() error {
	return c.categoryRepo.UpdateCategory(types.B{"_id": c.categoryId}, types.B{"stockNum": c.categoryEn.StockNum})
}

// existCategoryByName 分类名称对应分类是否存在
func (c *InventoryAggregate) existCategoryByName(name string) bool {
	if cg, _ := c.categoryRepo.GetCategoryByCategoryName(name); cg != nil && cg.Id != constants.EmptyStr {
		return true
	} else {
		return false
	}
}

// existCategoryById 分类id对应分类是否存在
func (c *InventoryAggregate) existCategoryById(id string) bool {
	if en, _ := c.categoryRepo.GetCategoryById(id); en != nil && en.Id != constants.EmptyStr {
		return true
	} else {
		return false
	}
}
