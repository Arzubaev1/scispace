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

type reportRepo struct {
	db *pgxpool.Pool
}

func NewReportRepo(db *pgxpool.Pool) *reportRepo {
	return &reportRepo{
		db: db,
	}
}

func (r *reportRepo) Create(ctx context.Context, req *models.CreateReport) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO reports(id, report_status, description,updated_at)
		VALUES ($1, $2, $3, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.ReportStatus,
		req.Description,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *reportRepo) GetByID(ctx context.Context, req *models.ReportPrimaryKey) (*models.Report, error) {

	// var whereField = "id"
	// if len(req.Email) > 0 {
	// 	whereField = "email"
	// 	req.Id = req.Email
	// }

	var (
		query string

		id            sql.NullString
		report_status sql.NullString
		description   sql.NullString
		createdAt     sql.NullString
		updatedAt     sql.NullString
	)

	query = `
		SELECT
			id,
			report_status,
			description,
			created_at,
			updated_at
		FROM reports
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&report_status,
		&description,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Report{
		Id:           id.String,
		ReportStatus: report_status.String,
		Description:  description.String,
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}, nil
}

func (r *reportRepo) GetList(ctx context.Context, req *models.ReportGetListRequest) (*models.ReportGetListResponse, error) {

	var (
		resp   = &models.ReportGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			report_status,
			description,
			created_at,
			updated_at
		FROM reports
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND description ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id            sql.NullString
			report_status sql.NullString
			description   sql.NullString
			createdAt     sql.NullString
			updatedAt     sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&report_status,
			&description,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Reports = append(resp.Reports, &models.Report{
			Id:           id.String,
			ReportStatus: report_status.String,
			Description:  description.String,
			CreatedAt:    createdAt.String,
			UpdatedAt:    updatedAt.String,
		})
	}

	return resp, nil
}

func (r *reportRepo) Update(ctx context.Context, req *models.UpdateReport) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		"reports"
		SET
		id = :id,
		report_status = :report_status,
		description = :description,
			updated_at = NOW()
		  WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"report_status": req.ReportStatus,
		"description":   req.Description,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *reportRepo) Delete(ctx context.Context, req *models.ReportPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM reports WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
