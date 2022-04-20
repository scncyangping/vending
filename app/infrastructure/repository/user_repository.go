package repository

import (
	"vending/app/domain/entity"
	"vending/app/domain/vo"
	"vending/app/infrastructure/converter"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
	"vending/app/infrastructure/pkg/log"
	"vending/app/types"
)

type UserRepository struct {
	mgo *mongo.MgoV
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		mgo: mongo.OpCn("user"),
	}
}

func (u *UserRepository) CreateUser(entity *entity.UserEntity) string {
	userDo := converter.ConvertUserE2UserD(entity)
	return u.mgo.InsertOne(userDo)
}

func (u *UserRepository) DeleteUser(s string) error {
	u.mgo.DeleteOne(types.B{"_id": s})
	return nil
}

func (u *UserRepository) GetUserById(s string) *vo.UserVo {
	var (
		err  error
		uv   *vo.UserVo
		user do.UserDo
	)
	err = u.mgo.FindOne(types.B{"_id": s}, &user)
	if err != nil {
		log.Logger().Error("GetUserById Error, %v", err)
	} else {
		uv = converter.ConvertUserD2UserV(&user)
	}
	return uv
}

func (u *UserRepository) ListUserBy(m map[string]interface{}) []*vo.UserVo {
	var (
		err   error
		users []*do.UserDo
		vos   []*vo.UserVo
	)
	err = u.mgo.Find(m, &users)
	if err != nil {
		log.Logger().Error("GetUserBy Error, %v", m)
	} else {
		for _, v := range users {
			vos = append(vos, converter.ConvertUserD2UserV(v))
		}
	}
	return vos
}

func (u *UserRepository) ListUserPageBy(skip, limit int64, sort, filter interface{}) []*vo.UserVo {
	var (
		err   error
		users []*do.UserDo
		vos   []*vo.UserVo
	)
	err = u.mgo.FindBy(skip, limit, sort, filter, &users)
	if err != nil {
		log.Logger().Error("GetUserBy Error, %v", err)
	} else {
		for _, v := range users {
			vos = append(vos, converter.ConvertUserD2UserV(v))
		}
	}
	return vos
}
