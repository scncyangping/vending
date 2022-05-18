package service

import (
	"vending/app/domain/dto"
	"vending/app/domain/entity"
	"vending/app/domain/service/imp/auth"
	"vending/app/domain/vo"
	"vending/app/infrastructure/repository"
)

type AuthService interface {
	LoginByName(name, pwd string) (*vo.UserVo, error)
	Register(*entity.UserEn) (string, error)
}

type CategoryService interface {
	SaveCategory(req *dto.CategorySaveReq) (string, error)
}

type Service struct {
	UserSrv *auth.UserServiceImpl
}

// NewService wire
func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserSrv: &auth.UserServiceImpl{
			UserRepo: repo.UserRepo,
		},
	}
}
