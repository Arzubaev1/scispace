package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type store struct {
	db            *pgxpool.Pool
	user          *userRepo
	question      *questionRepo
	post          *postRepo
	report        *reportRepo
	tool          *toolRepo
	database      *databaseRepo
	oqituvchi     *oqituvchiRepo
	tadqiqotchi   *tadqiqotchiRepo
	other         *otherRepo
	mutahassislik *mutahassislikRepo
	fan_tarmogi   *fan_tarmogiRepo
	mavzu         *mavzuRepo
	ish_joyi      *ish_joyiRepo
	maqola        *maqolaRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

func (s *store) Question() storage.QuestionRepoI {

	if s.question == nil {
		s.question = NewQuestionRepo(s.db)
	}

	return s.question
}

func (s *store) Post() storage.PostRepoI {

	if s.post == nil {
		s.post = NewPostRepo(s.db)
	}

	return s.post
}

func (s *store) Report() storage.ReportRepoI {

	if s.report == nil {
		s.report = NewReportRepo(s.db)
	}

	return s.report
}

func (s *store) Tool() storage.ToolRepoI {

	if s.tool == nil {
		s.tool = NewToolRepo(s.db)
	}

	return s.tool
}

func (s *store) Database() storage.DatabaseRepoI {

	if s.database == nil {
		s.database = NewDatabaseRepo(s.db)
	}

	return s.database
}
func (s *store) Oqituvchi() storage.OqituvchiRepoI {
	if s.oqituvchi == nil {
		s.oqituvchi = NewOqituvchiRepo(s.db)
	}
	return s.oqituvchi
}

func (s *store) Tadqiqotchi() storage.TadqiqotchiRepoI {
	if s.tadqiqotchi == nil {
		s.tadqiqotchi = NewTadqiqotchiRepo(s.db)
	}
	return s.tadqiqotchi
}
func (s *store) Other() storage.OtherRepoI {
	if s.other == nil {
		s.other = NewOtherRepo(s.db)
	}
	return s.other
}
func (s *store) IshJoyi() storage.IshJoyiRepoI {
	if s.ish_joyi == nil {
		s.ish_joyi = NewIshJoyiRepo(s.db)
	}
	return s.ish_joyi
}
func (s *store) Mutahassislik() storage.MutahassislikRepoI {
	if s.mutahassislik == nil {
		s.mutahassislik = NewMutahassislikRepo(s.db)
	}
	return s.mutahassislik
}
func (s *store) FanTarmogi() storage.FanTarmogiRepoI {
	if s.fan_tarmogi == nil {
		s.fan_tarmogi = NewFanTarmogiRepo(s.db)
	}
	return s.fan_tarmogi

}
func (s *store) Mavzu() storage.MavzuRepoI {
	if s.mavzu == nil {
		s.mavzu = NewMavzuRepo(s.db)
	}
	return s.mavzu
}
func (s *store) Maqola() storage.MaqolaRepoI {
	if s.maqola == nil {
		s.maqola = NewMaqolaRepo(s.db)
	}
	return s.maqola
}
