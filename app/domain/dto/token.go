package dto

type CreateTokenReq struct {
	Name string `json:"name"`
}

type CreateTokenRsp struct {
	Token string `json:"token"`
}
