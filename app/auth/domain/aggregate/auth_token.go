package aggregate

import (
	"context"
	"fmt"
	"strconv"
	"vending/app/auth/domain/dto"
	"vending/app/auth/domain/repo"
	"vending/app/auth/infrastructure/pkg/hcode"
	"vending/app/auth/infrastructure/pkg/tool"
)

type AuthToken struct {
	openId        string
	appId         string
	authTokenRepo repo.AuthTokenRepo
}

func (a AuthToken) GetUserInfo(ctx context.Context) (userSimple dto.UserSimple, err error) {
	var (
		uid int
	)
	uidByte, err := tool.AesECBDecrypt(fmt.Sprint(a.appId), a.openId)
	if err != nil {
		err = hcode.ServerErr
		return
	}
	uid, err = strconv.Atoi(string(uidByte))
	if err != nil {
		err = hcode.TranErr
		return
	}
	fmt.Println("uid", uid)
	// TODO 这里可以在 adpter 层去实现获取用户信息
	return dto.UserSimple{
		OpenId:   a.openId,
		Username: "",
		Phone:    "",
		Avatar:   "",
	}, nil
}
