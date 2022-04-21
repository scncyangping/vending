package entity

type RoleEntity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type uint8  `json:"type"`
}
