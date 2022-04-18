package entity

import "vending/app/infrastructure/pkg/util/snowflake"

type Id string

func NewId() Id {
	return Id(snowflake.NextId())
}
