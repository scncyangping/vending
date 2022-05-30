package inventory

import (
	"errors"
	"vending/app/application/cqe/cmd"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types"
	"vending/app/types/constants"
)

type DoInventorySrvImpl struct {
	categoryRepo repo.CategoryRepo
	stockRepo    repo.StockRepo
}

func NewDoInventorySrvImpl(categoryRepo repo.CategoryRepo, stockRepo repo.StockRepo) *DoInventorySrvImpl {
	return &DoInventorySrvImpl{categoryRepo: categoryRepo, stockRepo: stockRepo}
}

func (c *DoInventorySrvImpl) QueryCategoryPageBy(skip, limit int64, sort, filter any) ([]*do.CategoryDo, error) {
	return c.categoryRepo.ListCategoryPageBy(skip, limit, sort, filter)
}

func (c *DoInventorySrvImpl) QueryStockPageBy(skip, limit int64, sort, filter any) ([]*do.StockDo, error) {
	return c.stockRepo.ListStockPageBy(skip, limit, sort, filter)
}

// RemoveCategoryByIds 批量移除分类
func (c *DoInventorySrvImpl) RemoveCategoryByIds(ids []string) error {
	// 查询分类下是否有库存数据，若有，不允许移除
	if l, err := c.stockRepo.ListStockByIdsAndStatus(ids, types.StockNormal); err != nil {
		return err
	} else {
		if len(l) > 0 {
			return errors.New("该分类下存在未使用库存,不允许删除")
		}
	}
	return c.categoryRepo.DeleteCategoryByIds(ids)
}

// SaveCategory 添加分类
func (c *DoInventorySrvImpl) SaveCategory(req *cmd.CategorySaveCmd) (string, error) {
	var (
		err        error
		categoryEn entity.CategoryEn
	)
	// 查询对应分类是否存在
	if c.existCategoryByName(req.Name) {
		return constants.EmptyStr, errors.New("该分类已存在")
	}

	categoryEn.Name = req.Name
	categoryEn.PId = req.PId
	categoryEn.SellType = req.SellType
	categoryEn.Id = snowflake.NextId()

	if _, err = c.categoryRepo.SaveCategory(&categoryEn); err != nil {
		return constants.EmptyStr, err
	}
	return categoryEn.Id, nil
}

// existCategoryByName 分类名称对应分类是否存在
func (c *DoInventorySrvImpl) existCategoryByName(name string) bool {
	if cg, _ := c.categoryRepo.GetCategoryByCategoryName(name); cg != nil && cg.Id != constants.EmptyStr {
		return true
	} else {
		return false
	}
}
