package dto

type CategorySaveReq struct {
	Name string `json:"name" bson:"name"`
	PId  string `json:"pId" bson:"pId"`
}
