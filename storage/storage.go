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
	Oqituvchi() OqituvchiRepoI
	Tadqiqotchi() TadqiqotchiRepoI
	Other() OtherRepoI
	Mutahassislik() MutahassislikRepoI
	IshJoyi() IshJoyiRepoI
	Mavzu() MavzuRepoI
	FanTarmogi() FanTarmogiRepoI
	Maqola() MaqolaRepoI
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
type OqituvchiRepoI interface {
	Create(context.Context, *models.CreateOqituvchi) (string, error)
	GetByID(context.Context, *models.OqituvchiPrimaryKey) (*models.Oqituvchi, error)
	GetList(context.Context, *models.OqituvchiGetListRequest) (*models.OqituvchiGetListResponse, error)
	Update(context.Context, *models.UpdateOqituvchi) (int64, error)
	Delete(context.Context, *models.OqituvchiPrimaryKey) error
}

type TadqiqotchiRepoI interface {
	Create(context.Context, *models.CreateTadqiqotchi) (string, error)
	GetByID(context.Context, *models.TadqiqotchiPrimaryKey) (*models.Tadqiqotchi, error)
	GetList(context.Context, *models.TadqiqotchiGetListRequest) (*models.TadqiqotchiGetListResponse, error)
	Update(context.Context, *models.UpdateTadqiqotchi) (int64, error)
	Delete(context.Context, *models.TadqiqotchiPrimaryKey) error
}
type OtherRepoI interface {
	Create(context.Context, *models.CreateOther) (string, error)
	GetByID(context.Context, *models.OtherPrimaryKey) (*models.Other, error)
	GetList(context.Context, *models.OtherGetListRequest) (*models.OtherGetListResponse, error)
	Update(context.Context, *models.UpdateOther) (int64, error)
	Delete(context.Context, *models.OtherPrimaryKey) error
}
type MutahassislikRepoI interface {
	Create(context.Context, *models.CreateMutahassislik) (string, error)
	GetByID(context.Context, *models.MutahassislikPrimaryKey) (*models.Mutahassislik, error)
	GetList(context.Context, *models.MutahassislikGetListRequest) (*models.MutahassislikGetListResponse, error)
	Update(context.Context, *models.UpdateMutahassislik) (int64, error)
	Delete(context.Context, *models.MutahassislikPrimaryKey) error
}
type IshJoyiRepoI interface {
	Create(context.Context, *models.CreateIshJoyi) (string, error)
	GetByID(context.Context, *models.IshJoyiPrimaryKey) (*models.IshJoyi, error)
	GetList(context.Context, *models.IshJoyiGetListRequest) (*models.IshJoyiGetListResponse, error)
	Update(context.Context, *models.UpdateIshJoyi) (int64, error)
	Delete(context.Context, *models.IshJoyiPrimaryKey) error
}
type MavzuRepoI interface {
	Create(context.Context, *models.CreateMavzu) (string, error)
	GetByID(context.Context, *models.MavzuPrimaryKey) (*models.Mavzu, error)
	GetList(context.Context, *models.MavzuGetListRequest) (*models.MavzuGetListResponse, error)
	Update(context.Context, *models.UpdateMavzu) (int64, error)
	Delete(context.Context, *models.MavzuPrimaryKey) error
}
type FanTarmogiRepoI interface {
	Create(context.Context, *models.CreateFanTarmogi) (string, error)
	GetByID(context.Context, *models.FanTarmogiPrimaryKey) (*models.FanTarmogi, error)
	GetList(context.Context, *models.FanTarmogiGetListRequest) (*models.FanTarmogiGetListResponse, error)
	Update(context.Context, *models.UpdateFanTarmogi) (int64, error)
	Delete(context.Context, *models.FanTarmogiPrimaryKey) error
}
type MaqolaRepoI interface {
	Create(context.Context, *models.CreateMaqola) (string, error)
	GetByID(context.Context, *models.MaqolaPrimaryKey) (*models.Maqola, error)
	GetList(context.Context, *models.MaqolaGetListRequest) (*models.MaqolaGetListResponse, error)
	Update(context.Context, *models.UpdateMaqola) (int64, error)
	Delete(context.Context, *models.MaqolaPrimaryKey) error
}
