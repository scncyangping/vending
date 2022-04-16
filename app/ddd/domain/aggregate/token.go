package aggregate

import (
	"vending/app/ddd/domain/entity"
)

type Token struct {
	Id   entity.Id
	Name string
	// 添加其他待插入数据
}
