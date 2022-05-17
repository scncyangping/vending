package mgoRepo

import (
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/converter"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/database/mongo"
	"vending/app/infrastructure/pkg/log"
	"vending/app/types"
)

var _ repo.UserRepo = (*UserMgoRepository)(nil)

type UserMgoRepository struct {
	mgo *mongo.MgoV
}

func NewUserRepository(m *mongo.MgoV) *UserMgoRepository {
	return &UserMgoRepository{mgo: m}
}

func (u *UserMgoRepository) SaveUser(entity *entity.UserEn) (string, error) {
	userDo := converter.ConvertUserE2UserD(entity)
	return u.mgo.InsertOne(userDo)
}

func (u *UserMgoRepository) UpdateUser(filter types.B, update types.B) error {
	if _, err := u.mgo.Update(filter, update); err != nil {
		log.Logger().Error("UpdateUser Error, %v", err)
		return err
	}
	return nil
}

func (u *UserMgoRepository) DeleteUser(s string) error {
	u.mgo.DeleteOne(types.B{"_id": s})
	return nil
}

func (u *UserMgoRepository) GetUserByName(name string) (*do.UserDo, error) {
	var (
		err   error
		users do.UserDo
	)
	err = u.mgo.FindOne(types.B{"name": name}, &users)
	if err != nil {
		log.Logger().Error("GetUserByName Error, %v", err)
		return nil, err
	}
	return &users, nil
}

func (u *UserMgoRepository) GetUserById(s string) (*do.UserDo, error) {
	var (
		err   error
		users do.UserDo
	)
	if err = u.mgo.FindOne(types.B{"_id": s}, &users); err != nil {
		log.Logger().Error("GetUserById Error, %v", err)
		return nil, err
	}
	return &users, nil
}

func (u *UserMgoRepository) ListUserBy(m map[string]interface{}) ([]*do.UserDo, error) {
	var (
		err   error
		users []*do.UserDo
	)
	if err = u.mgo.Find(m, &users); err != nil {
		log.Logger().Error("GetUserBy Error, %v", m)
		return nil, err
	}
	return users, nil
}

func (u *UserMgoRepository) ListUserPageBy(skip, limit int64, sort, filter interface{}) ([]*do.UserDo, error) {
	var (
		err   error
		users []*do.UserDo
	)
	if err = u.mgo.FindBy(skip, limit, sort, filter, &users); err != nil {
		log.Logger().Error("GetUserBy Error, %v", err)
		return nil, err
	}
	return users, nil
}
