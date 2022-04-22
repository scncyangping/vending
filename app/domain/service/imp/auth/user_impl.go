package auth

import (
	"errors"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/domain/vo"
	"vending/app/infrastructure/pkg/util"
	"vending/app/infrastructure/pkg/util/jwt"
	"vending/app/types/constants"
)

type UserServiceImpl struct {
	UserRepo repo.UserRepo
}

// Register 注册用户
func (a *UserServiceImpl) Register(rr *entity.UserEn) (string, error) {
	// 校验用户是否存在 是否有其他关联条件
	if userDo, err := a.UserRepo.GetUserByName(rr.Name); err != nil {
		return constants.EmptyStr, err
	} else {
		if userDo != nil {
			return constants.EmptyStr, errors.New("user is exist")
		}
	}
	return a.UserRepo.SaveUser(rr)
}

func (a *UserServiceImpl) LoginByName(name, pwd string) (*vo.UserVo, error) {
	var (
		rp vo.UserVo
	)
	// 查询数据
	if userDo, err := a.UserRepo.GetUserByName(name); err != nil {
		return nil, err
	} else {
		if userDo == nil {
			err = errors.New("user name or pwd error")
		} else {
			if userDo.Pwd != pwd {
				err = errors.New("user name or pwd error")
			}
		}
		if err != nil {
			return nil, err
		} else {
			if token, err := jwt.GenerateToken(userDo.Name); err != nil {
				return nil, err
			} else {
				rp.Token = token
			}
		}
		util.StructCopy(&rp, userDo)

		return &rp, nil
	}
}
