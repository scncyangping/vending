package entity

import "vending/app/ddd/infrastructure/pkg/util/snowflake"

type Id string

func NewId() Id {
	return Id(snowflake.NextId())
}
