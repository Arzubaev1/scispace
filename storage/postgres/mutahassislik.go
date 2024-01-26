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

type mutahassislikRepo struct {
	db *pgxpool.Pool
}

func NewMutahassislikRepo(db *pgxpool.Pool) *mutahassislikRepo {
	return &mutahassislikRepo{
		db: db,
	}
}

func (r *mutahassislikRepo) Create(ctx context.Context, req *models.CreateMutahassislik) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO mutahassislik(
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

func (r *mutahassislikRepo) GetByID(ctx context.Context, req *models.MutahassislikPrimaryKey) (*models.Mutahassislik, error) {

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
		FROM mutahassislik
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

	return &models.Mutahassislik{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (r *mutahassislikRepo) GetList(ctx context.Context, req *models.MutahassislikGetListRequest) (*models.MutahassislikGetListResponse, error) {

	var (
		resp   = &models.MutahassislikGetListResponse{}
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
			
		FROM mutahassislik
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

		resp.Mutahassisliklar = append(resp.Mutahassisliklar, &models.Mutahassislik{
			Id:        id.String,
			Name:      name.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}

	return resp, nil
}

func (r *mutahassislikRepo) Update(ctx context.Context, req *models.UpdateMutahassislik) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			mutahassislik
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

func (r *mutahassislikRepo) Delete(ctx context.Context, req *models.MutahassislikPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM mutahassislik WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
