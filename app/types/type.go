package types

import (
	"go.mongodb.org/mongo-driver/bson"
)

type AuthenticationType string

const (
	JWT AuthenticationType = "JWT"
)

type ResultCode string
type ResultMsg string

type UserType uint8
type UserStatus uint8

type CategoryStatus uint8

const (
	ADMIN UserType = 1 << iota
	USER

	NORMAL UserStatus = 1 << iota
	FROZEN

	ON CategoryStatus = 1 << iota
	OFF
)

type B bson.M
