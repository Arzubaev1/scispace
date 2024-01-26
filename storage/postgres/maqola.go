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

type maqolaRepo struct {
	db *pgxpool.Pool
}

func NewMaqolaRepo(db *pgxpool.Pool) *maqolaRepo {
	return &maqolaRepo{
		db: db,
	}
}

// FullName    string `json:"fullname"`
// 	Institution string `json:"institution"`
// 	Department  string `json:"department"`
// 	Degree      string `json:"degree"`
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`

func (r *maqolaRepo) Create(ctx context.Context, req *models.CreateMaqola) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO maqola(
			id,
			name,
			tavsifi,
			qoshimcha_linklar,
			updated_at)
		VALUES ($1, $2, $3, $4, NOW())
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

func (r *maqolaRepo) GetByID(ctx context.Context, req *models.MaqolaPrimaryKey) (*models.Maqola, error) {

	var (
		query             string
		id                sql.NullString
		name              sql.NullString
		tavsifi           sql.NullString
		qoshimcha_linklar sql.NullString
		created_at        sql.NullString
		updated_at        sql.NullString
	)

	query = `
		SELECT
			id
			name,
			tavsifi,
			qoshimcha_linklar,
			created_at,
			updated_at
		FROM maqola
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&tavsifi,
		&qoshimcha_linklar,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Maqola{
		Id:               id.String,
		Name:             name.String,
		Tavsifi:          tavsifi.String,
		QoshimchaLinklar: qoshimcha_linklar.String,
		CreatedAt:        created_at.String,
		UpdatedAt:        updated_at.String,
	}, nil
}

func (r *maqolaRepo) GetList(ctx context.Context, req *models.MaqolaGetListRequest) (*models.MaqolaGetListResponse, error) {

	var (
		resp   = &models.MaqolaGetListResponse{}
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
			tavsifi,
			qoshimcha_linklar,
			created_at,
			updated_at
			
		FROM maqola
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id                sql.NullString
			name              sql.NullString
			tavsifi           sql.NullString
			qoshimcha_linklar sql.NullString
			created_at        sql.NullString
			updated_at        sql.NullString
		)

		err := rows.Scan(
			&id,
			&name,
			&tavsifi,
			&qoshimcha_linklar,
			&created_at,
			&updated_at,
		)

		if err != nil {
			return nil, err
		}

		resp.Maqolalar = append(resp.Maqolalar, &models.Maqola{
			Id:               id.String,
			Name:             name.String,
			Tavsifi:          tavsifi.String,
			QoshimchaLinklar: qoshimcha_linklar.String,
			CreatedAt:        created_at.String,
			UpdatedAt:        updated_at.String,
		})
	}

	return resp, nil
}

func (r *maqolaRepo) Update(ctx context.Context, req *models.UpdateMaqola) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			maqola
		SET
		id = :id,
		name = :name,
		tavsifi = :tavsifi,
		qoshimcha_linklar = :qoshimcha_linklar,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":                req.Id,
		"name":              req.Name,
		"tavsifi":           req.Tavsifi,
		"qoshimcha_linklar": req.QoshimchaLinklar,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *maqolaRepo) Delete(ctx context.Context, req *models.MaqolaPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM maqola WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
