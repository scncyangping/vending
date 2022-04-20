package do

type UserDo struct {
	Id         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	NickName   string `json:"nickName" bson:"nickName"`
	Phone      string `json:"phone" bson:"phone"`
	Email      string `json:"email" bson:"email"`
	Pwd        string `json:"pwd" bson:"pwd"`
	Type       uint8  `json:"type" bson:"type"`
	Status     uint8  `json:"status"bson:"status"`
	CreateTime int64  `json:"createTime" bson:"createTime"`
	UpdateTime int64  `json:"updateTime" bson:"updateTime"`
	IsDeleted  uint8  `json:"isDeleted" bson:"isDeleted"`
}
