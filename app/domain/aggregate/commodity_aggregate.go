package aggregate

import "vending/app/domain/dto"

type CommodityAg struct {
	Id            string            `json:"id"`
	Name          string            `json:"name"`
	Des           string            `json:"des"`
	Introduction  string            `json:"introduction"`
	Type          uint8             `json:"type"`
	CommodityData dto.CommodityData `json:"commodityData"`
}

type CommodityAggregate struct {
}
