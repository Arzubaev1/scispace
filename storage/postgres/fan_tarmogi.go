package postgres

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type fan_tarmogiRepo struct {
	db *pgxpool.Pool
}

func NewFanTarmogiRepo(db *pgxpool.Pool) *fan_tarmogiRepo {
	return &fan_tarmogiRepo{
		db: db,
	}
}

// FullName    string `json:"fullname"`
// 	Institution string `json:"institution"`
// 	Department  string `json:"department"`
// 	Degree      string `json:"degree"`
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`

func (r *fan_tarmogiRepo) Create(ctx context.Context, req *models.CreateFanTarmogi) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO fan_tarmogi(
			id,
			name,
			updated_at)
		VALUES ($1, $2, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *fan_tarmogiRepo) GetByID(ctx context.Context, req *models.FanTarmogiPrimaryKey) (*models.FanTarmogi, error) {

	var (
		query      string
		id         sql.NullString
		name       sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	query = `
		SELECT
			id
			name,
			created_at,
			updated_at
		FROM fan_tarmogi
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.FanTarmogi{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (r *fan_tarmogiRepo) GetList(ctx context.Context, req *models.FanTarmogiGetListRequest) (*models.FanTarmogiGetListResponse, error) {

	var (
		resp   = &models.FanTarmogiGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			created_at,
			updated_at
			
		FROM fan_tarmogi
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND first_name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id         sql.NullString
			name       sql.NullString
			created_at sql.NullString
			updated_at sql.NullString
		)

		err := rows.Scan(
			&id,
			&name,
			&created_at,
			&updated_at,
		)

		if err != nil {
			return nil, err
		}

		resp.FanTarmoqlari = append(resp.FanTarmoqlari, &models.FanTarmogi{
			Id:        id.String,
			Name:      name.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}

	return resp, nil
}

func (r *fan_tarmogiRepo) Update(ctx context.Context, req *models.UpdateFanTarmogi) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			fan_tarmogi
		SET
		id = :id,
		name = :name,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":   req.Id,
		"name": req.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *fan_tarmogiRepo) Delete(ctx context.Context, req *models.FanTarmogiPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM fan_tarmogi WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
