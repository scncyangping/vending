package handler

type Base struct {
	CreateTime    int64 `bson:"createTime" json:"createTime"`
	UpdateTime    int64 `bson:"updateTime" json:"updateTime"`
	LastLoginTime int64 `bson:"lastLoginTime" json:"lastLoginTime"`
	IsDeleted     bool  `bson:"isDeleted" json:"isDeleted"` // 0 正常 1 删除
}
