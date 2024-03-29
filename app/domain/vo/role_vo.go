package vo

import "vending/app/domain/obj"

type RoleVo struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	Type        uint8            `json:"type"`
	Permissions []obj.Permission `json:"permission"`
}
