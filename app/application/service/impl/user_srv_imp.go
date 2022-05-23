package impl

import (
	"vending/app/application/cqe/cmd"
	"vending/app/application/dto"
	"vending/app/domain/entity"
	"vending/app/domain/service"
	"vending/app/infrastructure/pkg/util"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types/constants"
)

// 不要处理业务逻辑
// ：参数验证、错误处理、监控日志、事务处理、认证与授权

type AuthSrvImp struct {
	authSrv service.AuthService
}

// NewAuthSrvImp wire
func NewAuthSrvImp(authService *service.Service) *AuthSrvImp {
	return &AuthSrvImp{
		authSrv: authService.UserSrv,
	}
}

func (a *AuthSrvImp) Login(re cmd.LoginCmd) (dto.UserDto, error) {
	var (
		userDto dto.UserDto
	)
	if vo, err := a.authSrv.LoginByName(re.Name, re.Pwd); err != nil {
		return userDto, err
	} else {
		util.StructCopy(&userDto, vo)
		return userDto, nil
	}
}

func (a *AuthSrvImp) Register(re cmd.RegisterCmd) (string, error) {
	var ue entity.UserEn

	util.StructCopy(ue, re)

	if ue.Id == constants.EmptyStr {
		ue.Id = snowflake.NextId()
	}
	//if ue.Type == constants.ZERO {
	//	ue.Status = types.NORMAL
	//}
	return a.authSrv.Register(&ue)
}
