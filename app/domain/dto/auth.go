package dto

import (
	"github.com/dgrijalva/jwt-go"
	"vending/app/types"
)

type LoginRe struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type LoginRp struct {
	Token string `json:"token"`
}

type JwtToken struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}

type RegisterRe struct {
	Name     string         `json:"name"`
	NickName string         `json:"nickName"`
	Phone    string         `json:"phone"`
	Email    string         `json:"email"`
	Pwd      string         `json:"pwd"`
	Type     types.UserType `json:"type" bson:"type"`
}
