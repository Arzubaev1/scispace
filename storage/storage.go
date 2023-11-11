package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	User() UserRepoI
	Question() QuestionRepoI
	Post() PostRepoI
	Report() ReportRepoI
	Tool() ToolRepoI
	Database() DatabaseRepoI
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

type PostRepoI interface {
	Create(context.Context, *models.CreatePost) (string, error)
	GetByID(context.Context, *models.PostPrimaryKey) (*models.Post, error)
	GetList(context.Context, *models.PostGetListRequest) (*models.PostGetListResponse, error)
	Update(context.Context, *models.UpdatePost) (int64, error)
	Delete(context.Context, *models.PostPrimaryKey) error
}

type ReportRepoI interface {
	Create(context.Context, *models.CreateReport) (string, error)
	GetByID(context.Context, *models.ReportPrimaryKey) (*models.Report, error)
	GetList(context.Context, *models.ReportGetListRequest) (*models.ReportGetListResponse, error)
	Update(context.Context, *models.UpdateReport) (int64, error)
	Delete(context.Context, *models.ReportPrimaryKey) error
}

type ToolRepoI interface {
	Create(context.Context, *models.CreateTool) (string, error)
	GetByID(context.Context, *models.ToolPrimaryKey) (*models.Tool, error)
	GetList(context.Context, *models.ToolGetListRequest) (*models.ToolGetListResponse, error)
	Update(context.Context, *models.UpdateTool) (int64, error)
	Delete(context.Context, *models.ToolPrimaryKey) error
}

type DatabaseRepoI interface {
	Create(context.Context, *models.CreateDatabase) (string, error)
	GetByID(context.Context, *models.DatabasePrimaryKey) (*models.Database, error)
	GetList(context.Context, *models.DatabaseGetListRequest) (*models.DatabaseGetListResponse, error)
	Update(context.Context, *models.UpdateDatabase) (int64, error)
	Delete(context.Context, *models.DatabasePrimaryKey) error
}
