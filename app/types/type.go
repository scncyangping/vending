package types

import (
	"go.mongodb.org/mongo-driver/bson"
)

type B bson.M

type AuthenticationType string

const (
	JWT AuthenticationType = "JWT"
)
