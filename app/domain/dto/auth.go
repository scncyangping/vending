package dto

type JwtAuthRe struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type JwtAuthRp struct {
	Token string `json:"token"`
}

type JwtAuthTokenRe struct {
	UserName string `json:"userName"`
}
