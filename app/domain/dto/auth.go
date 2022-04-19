package dto

type LoginRe struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type LoginRp struct {
	Token string `json:"token"`
}

type JwtAuthTokenRe struct {
	UserName string `json:"userName"`
}
