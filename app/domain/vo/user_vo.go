package vo

type UserVo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Type       uint8  `json:"type"`
	Status     uint8  `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Token      string `json:"token"`
}
