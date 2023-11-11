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

type questionRepo struct {
	db *pgxpool.Pool
}

func NewQuestionRepo(db *pgxpool.Pool) *questionRepo {
	return &questionRepo{
		db: db,
	}
}

func (r *questionRepo) Create(ctx context.Context, req *models.CreateQuestion) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO questions(id, title, context,tags,updated_at)
		VALUES ($1, $2, $3, $4, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Title,
		req.Context,
		req.Tags,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *questionRepo) GetByID(ctx context.Context, req *models.QuestionPrimaryKey) (*models.Question, error) {

	// var whereField = "id"
	// if len(req.Email) > 0 {
	// 	whereField = "email"
	// 	req.Id = req.Email
	// }

	var (
		query string

		id        sql.NullString
		title     sql.NullString
		context   sql.NullString
		tags      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT
			id,
			title,
			context,
			tags,
			created_at,
			updated_at
		FROM questions
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&title,
		&context,
		&tags,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Question{
		Id:        id.String,
		Title:     title.String,
		Context:   context.String,
		Tags:      tags.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *questionRepo) GetList(ctx context.Context, req *models.QuestionGetListRequest) (*models.QuestionGetListResponse, error) {

	var (
		resp   = &models.QuestionGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			title,
			context,
			tags,
			created_at,
			updated_at
		FROM questions
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND tags ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			title     sql.NullString
			context   sql.NullString
			tags      sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&title,
			&context,
			&tags,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Questions = append(resp.Questions, &models.Question{
			Id:        id.String,
			Title:     title.String,
			Context:   context.String,
			Tags:      tags.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *questionRepo) Update(ctx context.Context, req *models.UpdateQuestion) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	

	query = `
		UPDATE
		"questions"
		SET
		id = :id,
		title = :title,
		context = :context,
		tags = :tags,
			updated_at = NOW()
		  WHERE id = :id
	`

	params = map[string]interface{}{
		"id":      req.Id,
		"title":   req.Title,
		"context": req.Context,
		"tags":    req.Tags,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *questionRepo) Delete(ctx context.Context, req *models.QuestionPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM questions WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
