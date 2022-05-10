package dto

type CommoditySaveReq struct {
	Name string      `json:"name"`
	Type uint8       `json:"type"`
	Data interface{} `json:"data"`
}

type CommodityBatchSaveReq struct {
	List []CommoditySaveReq `json:"list"`
}
