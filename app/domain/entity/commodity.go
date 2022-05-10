package entity

type CommodityEn struct {
	Id         string      `json:"id" bson:"_id"`
	Name       string      `json:"name" bson:"name"`
	Type       uint8       `json:"type" bson:"type"`
	Data       interface{} `json:"data" bson:"data"`
	CreateTime int64       `json:"createTime" bson:"createTime"`
	UpdateTime int64       `json:"updateTime" bson:"updateTime"`
	IsDeleted  uint8       `json:"isDeleted" bson:"isDeleted"`
}
