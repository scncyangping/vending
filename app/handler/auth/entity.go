package auth

import (
	"vending/app/handler"
	"vending/common/util"
	"vending/common/util/snowflake"
)

type User struct {
	*handler.Base
	Id     string `bson:"_id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Phone  string `bson:"phone" json:"phone"`
	IdCard string `bson:"idCard" json:"idCard"`
	Pwd    string `bson:"pwd" json:"pwd"`
	Status int    `bson:"status" json:"status"` // 0 正常 1 冻结
}

func UserInstance(dto UserDTO) *User {
	return &User{
		Id:   snowflake.NextId(),
		Name: dto.Name,
		Pwd:  dto.Pwd,
		Base: &handler.Base{
			CreateTime: util.NowTimestamp(),
		},
	}
}

type UserDTO struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type UserVO struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	IdCard string `json:"idCard"`
	Status int    `json:"status"`
}
