package imp

//type CategoryServiceImpl struct {
//	categoryRepo repo.CategoryRepo
//}
//
//// SaveCategory 添加分类
//func (c *CategoryServiceImpl) SaveCategory(req *dto.CategorySaveReq) (string, error) {
//	var (
//		err        error
//		categoryEn entity.CategoryEn
//	)
//	// 查询对应分类是否存在
//	if c.existCategoryByName(req.Name) {
//		return constants.EmptyStr, errors.New("该分类已存在")
//	}
//
//	categoryEn.Name = req.Name
//	categoryEn.PId = req.PId
//	categoryEn.SellType = req.SellType
//	categoryEn.Id = snowflake.NextId()
//
//	if _, err = c.categoryRepo.SaveCategory(&categoryEn); err != nil {
//		return constants.EmptyStr, err
//	}
//	return categoryEn.Id, nil
//}
//
//// existCategoryByName 分类名称对应分类是否存在
//func (c *CategoryServiceImpl) existCategoryByName(name string) bool {
//	if cg, _ := c.categoryRepo.GetCategoryByCategoryName(name); cg != nil && cg.Id != constants.EmptyStr {
//		return true
//	} else {
//		return false
//	}
//}
