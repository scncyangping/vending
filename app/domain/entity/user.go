package entity

type UserEntity struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
	Type     uint8  `json:"type"`
	Status   uint8  `json:"status"`
}
