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

type toolRepo struct {
	db *pgxpool.Pool
}

func NewToolRepo(db *pgxpool.Pool) *toolRepo {
	return &toolRepo{
		db: db,
	}
}

func (r *toolRepo) Create(ctx context.Context, req *models.CreateTool) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO tools(id, tool_name, link,updated_at)
		VALUES ($1, $2, $3, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.ToolName,
		req.Link,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *toolRepo) GetByID(ctx context.Context, req *models.ToolPrimaryKey) (*models.Tool, error) {

	// var whereField = "id"
	// if len(req.Email) > 0 {
	// 	whereField = "email"
	// 	req.Id = req.Email
	// }

	var (
		query string

		id        sql.NullString
		tool_name sql.NullString
		link      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT
			id,
			tool_name,
			link,
			created_at,
			updated_at
		FROM tools
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&tool_name,
		&link,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Tool{
		Id:        id.String,
		ToolName:  tool_name.String,
		Link:      link.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *toolRepo) GetList(ctx context.Context, req *models.ToolGetListRequest) (*models.ToolGetListResponse, error) {

	var (
		resp   = &models.ToolGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			tool_name,
			link,
			created_at,
			updated_at
		FROM tools
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND tool_name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			tool_name sql.NullString
			link      sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&tool_name,
			&link,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Tools = append(resp.Tools, &models.Tool{
			Id:        id.String,
			ToolName:  tool_name.String,
			Link:      link.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *toolRepo) Update(ctx context.Context, req *models.UpdateTool) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		"tools"
		SET
		id = :id,
		tool_name = :tool_name,
		link = :link,
			updated_at = NOW()
		  WHERE id = :id
	`

	params = map[string]interface{}{
		"id":        req.Id,
		"tool_name": req.ToolName,
		"link":      req.Link,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *toolRepo) Delete(ctx context.Context, req *models.ToolPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM tools WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
