package repository

import (
	"hub_op/app/repository/mgoRepo"
	"hub_op/common/pkg/database/mongo"
)

type Repository struct {
	UserRepo       *mgoRepo.UserMgoRepository
	ArticleRepo    *mgoRepo.ArticleMgoRepository
	ProjectRepo    *mgoRepo.ProjectMgoRepository
	RepoRepo       *mgoRepo.RepoMgoRepository
	GitAccountRepo *mgoRepo.GitAccountMgoRepository
}

// NewRepository wire
func NewRepository() *Repository {
	return &Repository{
		UserRepo:       mgoRepo.NewUserRepository(mongo.OpCn("user")),
		ArticleRepo:    mgoRepo.NewArticleRepository(mongo.OpCn("article")),
		ProjectRepo:    mgoRepo.NewProjectRepository(mongo.OpCn("project")),
		RepoRepo:       mgoRepo.NewRepoRepository(mongo.OpCn("repo")),
		GitAccountRepo: mgoRepo.NewGitAccountRepository(mongo.OpCn("git_account")),
	}
}
