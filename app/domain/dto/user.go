package dto

type UserRegisterRq struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
	Type     uint8  `json:"type"`
	Status   uint8  `json:"status"`
}
