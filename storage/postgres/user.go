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

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

// FullName    string `json:"fullname"`
// 	Institution string `json:"institution"`
// 	Department  string `json:"department"`
// 	Degree      string `json:"degree"`
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO users(id, fullname, institution,department, degree,email,password,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.FullName,
		req.Institution,
		req.Department,
		req.Degree,
		req.Email,
		req.Password,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var whereField = "id"
	if len(req.Email) > 0 {
		whereField = "email"
		req.Id = req.Email
	}

	var (
		query string

		id          sql.NullString
		fullname    sql.NullString
		institution sql.NullString
		department  sql.NullString
		degree      sql.NullString
		email       sql.NullString
		password    sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		SELECT
			id,
			fullname,
			institution,
			department,
			degree,
			email,
			password,
			created_at,
			updated_at
		FROM users
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&fullname,
		&institution,
		&department,
		&degree,
		&email,
		&password,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:          id.String,
		FullName:    fullname.String,
		Institution: institution.String,
		Department:  department.String,
		Degree:      degree.String,
		Email:       email.String,
		Password:    password.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var (
		resp   = &models.UserGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			fullname,
			institution,
			department,
			degree,
			email,
			password,
			created_at,
			updated_at
		FROM users
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND fullname ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			fullname    sql.NullString
			institution sql.NullString
			department  sql.NullString
			degree      sql.NullString
			email       sql.NullString
			password    sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&fullname,
			&institution,
			&department,
			&degree,
			&email,
			&password,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:          id.String,
			FullName:    fullname.String,
			Institution: institution.String,
			Department:  department.String,
			Degree:      degree.String,
			Email:       email.String,
			Password:    password.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			users
		SET
		id = :id,
		fullname = :fullname,
		institution = :institution,
		department = :department,
		degree = :degree,
		email = :email,
		password = :password,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"fullname":    req.FullName,
		"institution": req.Institution,
		"department":  req.Department,
		"degree":      req.Degree,
		"email":       req.Email,
		"password":    req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
