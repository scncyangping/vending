package types

import (
	"go.mongodb.org/mongo-driver/bson"
)

type AuthenticationType string

const (
	JWT AuthenticationType = "JWT"
)

type ResultCode string // 返回代码
type ResultMsg string  // 返回信息

type UserType uint8   // 用户类型
type UserStatus uint8 // 用户状态

type Status uint8

const (
	ADMIN    UserType = 1 << iota // 管理员
	MERCHANT                      // 商家
	USER

	NORMAL UserStatus = 1 << iota
	FROZEN

	CREATE Status = 1 << iota
	PENDING
	ON
	OFF
)

type B bson.M
