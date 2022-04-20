package converter

import (
	"vending/app/domain/entity"
	"vending/app/domain/vo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/util"
)

func ConvertUserE2UserD(e *entity.UserEntity) (u *do.UserDo) {
	util.StructCopy(u, e)
	u.CreateTime = util.NowTimestamp()
	u.UpdateTime = util.NowTimestamp()
	return
}

func ConvertUserD2UserV(e *do.UserDo) (v *vo.UserVo) {
	util.StructCopy(v, e)
	v.CreateTime = util.TimeFormat(e.CreateTime)
	v.UpdateTime = util.TimeFormat(e.UpdateTime)
	return
}
