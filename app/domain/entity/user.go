package entity

import "vending/app/types"

type UserEn struct {
	Id       string           `json:"id"`
	Name     string           `json:"name"`
	NickName string           `json:"nickName"`
	Phone    string           `json:"phone"`
	Email    string           `json:"email"`
	Pwd      string           `json:"pwd"`
	Type     types.UserType   `json:"type"`
	Status   types.UserStatus `json:"status"`
}
