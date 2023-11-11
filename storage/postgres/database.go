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

type databaseRepo struct {
	db *pgxpool.Pool
}

func NewDatabaseRepo(db *pgxpool.Pool) *databaseRepo {
	return &databaseRepo{
		db: db,
	}
}

func (r *databaseRepo) Create(ctx context.Context, req *models.CreateDatabase) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO databases(id, database_name, link,updated_at)
		VALUES ($1, $2, $3, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.DatabaseName,
		req.Link,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *databaseRepo) GetByID(ctx context.Context, req *models.DatabasePrimaryKey) (*models.Database, error) {

	// var whereField = "id"
	// if len(req.Email) > 0 {
	// 	whereField = "email"
	// 	req.Id = req.Email
	// }

	var (
		query string

		id            sql.NullString
		database_name sql.NullString
		link          sql.NullString
		createdAt     sql.NullString
		updatedAt     sql.NullString
	)

	query = `
		SELECT
			id,
			database_name,
			link,
			created_at,
			updated_at
		FROM databases
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&database_name,
		&link,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Database{
		Id:           id.String,
		DatabaseName: database_name.String,
		Link:         link.String,
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}, nil
}

func (r *databaseRepo) GetList(ctx context.Context, req *models.DatabaseGetListRequest) (*models.DatabaseGetListResponse, error) {

	var (
		resp   = &models.DatabaseGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			database_name,
			link,
			created_at,
			updated_at
		FROM databases
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND database_name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id            sql.NullString
			database_name sql.NullString
			link          sql.NullString
			createdAt     sql.NullString
			updatedAt     sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&database_name,
			&link,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Databases = append(resp.Databases, &models.Database{
			Id:           id.String,
			DatabaseName: database_name.String,
			Link:         link.String,
			CreatedAt:    createdAt.String,
			UpdatedAt:    updatedAt.String,
		})
	}

	return resp, nil
}

func (r *databaseRepo) Update(ctx context.Context, req *models.UpdateDatabase) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		"databases"
		SET
		id = :id,
		database_name = :database_name,
		link = :link,
			updated_at = NOW()
		  WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"database_name": req.DatabaseName,
		"link":          req.Link,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *databaseRepo) Delete(ctx context.Context, req *models.DatabasePrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM databases WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
