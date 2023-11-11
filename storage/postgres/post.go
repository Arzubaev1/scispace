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

type postRepo struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *postRepo {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Create(ctx context.Context, req *models.CreatePost) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO posts(id, title, context,link,updated_at)
		VALUES ($1, $2, $3, $4, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Title,
		req.Context,
		req.Link,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *postRepo) GetByID(ctx context.Context, req *models.PostPrimaryKey) (*models.Post, error) {

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
		link      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT
			id,
			title,
			context,
			link,
			created_at,
			updated_at
		FROM posts
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&title,
		&context,
		&link,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Post{
		Id:        id.String,
		Title:     title.String,
		Context:   context.String,
		Link:      link.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *postRepo) GetList(ctx context.Context, req *models.PostGetListRequest) (*models.PostGetListResponse, error) {

	var (
		resp   = &models.PostGetListResponse{}
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
			link,
			created_at,
			updated_at
		FROM posts
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
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
			link      sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&title,
			&context,
			&link,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Posts = append(resp.Posts, &models.Post{
			Id:        id.String,
			Title:     title.String,
			Context:   context.String,
			Link:      link.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *postRepo) Update(ctx context.Context, req *models.UpdatePost) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		"posts"
		SET
		id = :id,
		title = :title,
		context = :context,
		link = :link,
			updated_at = NOW()
		  WHERE id = :id
	`

	params = map[string]interface{}{
		"id":      req.Id,
		"title":   req.Title,
		"context": req.Context,
		"link":    req.Link,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *postRepo) Delete(ctx context.Context, req *models.PostPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM posts WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
