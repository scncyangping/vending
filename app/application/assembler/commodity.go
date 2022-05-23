package assembler

import (
	"vending/app/application/dto"
	"vending/app/domain/entity"
	"vending/app/infrastructure/pkg/util"
)

func CommodityEnToDto(en *entity.CommodityEn) *dto.CommodityDto {
	dto := dto.CommodityDto{}
	util.StructCopy(&dto, en)
	return &dto
}
