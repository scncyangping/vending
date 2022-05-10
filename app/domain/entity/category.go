package entity

import "vending/app/types"

type CategoryEn struct {
	Id         string               `json:"id" bson:"_id"`
	Name       string               `json:"name" bson:"name"`
	PId        string               `json:"pId" bson:"pId"`
	Status     types.CategoryStatus `json:"status" bson:"status"`
	CreateTime int64                `json:"createTime" bson:"createTime"`
	UpdateTime int64                `json:"updateTime" bson:"updateTime"`
	IsDeleted  uint8                `json:"isDeleted" bson:"isDeleted"`
}
