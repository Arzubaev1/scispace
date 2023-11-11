package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	User() UserRepoI
	Question() QuestionRepoI
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
}

type QuestionRepoI interface {
	Create(context.Context, *models.CreateQuestion) (string, error)
	GetByID(context.Context, *models.QuestionPrimaryKey) (*models.Question, error)
	GetList(context.Context, *models.QuestionGetListRequest) (*models.QuestionGetListResponse, error)
	Update(context.Context, *models.UpdateQuestion) (int64, error)
	Delete(context.Context, *models.QuestionPrimaryKey) error
}
