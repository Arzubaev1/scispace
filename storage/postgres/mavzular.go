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

type mavzuRepo struct {
	db *pgxpool.Pool
}

func NewMavzuRepo(db *pgxpool.Pool) *mavzuRepo {
	return &mavzuRepo{
		db: db,
	}
}

// FullName    string `json:"fullname"`
// 	Institution string `json:"institution"`
// 	Department  string `json:"department"`
// 	Degree      string `json:"degree"`
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`

func (r *mavzuRepo) Create(ctx context.Context, req *models.CreateMavzu) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO mavzu(
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

func (r *mavzuRepo) GetByID(ctx context.Context, req *models.MavzuPrimaryKey) (*models.Mavzu, error) {

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
		FROM mavzu
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

	return &models.Mavzu{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (r *mavzuRepo) GetList(ctx context.Context, req *models.MavzuGetListRequest) (*models.MavzuGetListResponse, error) {

	var (
		resp   = &models.MavzuGetListResponse{}
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
			
		FROM mavzu
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

		resp.Mavzular = append(resp.Mavzular, &models.Mavzu{
			Id:        id.String,
			Name:      name.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}

	return resp, nil
}

func (r *mavzuRepo) Update(ctx context.Context, req *models.UpdateMavzu) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			mavzu
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

func (r *mavzuRepo) Delete(ctx context.Context, req *models.MavzuPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM mavzu WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
