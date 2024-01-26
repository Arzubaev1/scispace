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

type oqituvchiRepo struct {
	db *pgxpool.Pool
}

func NewOqituvchiRepo(db *pgxpool.Pool) *oqituvchiRepo {
	return &oqituvchiRepo{
		db: db,
	}
}

// FullName    string `json:"fullname"`
// 	Institution string `json:"institution"`
// 	Department  string `json:"department"`
// 	Degree      string `json:"degree"`
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`

func (r *oqituvchiRepo) Create(ctx context.Context, req *models.CreateOqituvchi) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO oqituvchi(
			id,
			first_name,
			last_name,
			middle_name,
			date_of_birth,
			ish_joyi,
			mutahassislik,
			fan_tarmogi,
			mavzular,
			email,
			password,
			phone_number,
			updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.LastName,
		req.MiddleName,
		req.DateOfBirth,
		req.IshJoyi,
		req.Mutahassislik,
		req.FanTarmogi,
		req.Mavzular,
		req.Email,
		req.Password,
		req.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *oqituvchiRepo) GetByID(ctx context.Context, req *models.OqituvchiPrimaryKey) (*models.Oqituvchi, error) {

	var whereField = "id"
	if len(req.Email) > 0 {
		whereField = "email"
		req.Id = req.Email
	}

	var (
		query         string
		id            sql.NullString
		first_name    sql.NullString
		last_name     sql.NullString
		middle_name   sql.NullString
		date_of_birth sql.NullString
		ish_joyi      sql.NullString
		mutahassislik sql.NullString
		fan_tarmogi   sql.NullString
		mavzular      sql.NullString
		email         sql.NullString
		password      sql.NullString
		phone_number  sql.NullString
		created_at    sql.NullString
		updated_at    sql.NullString
	)

	query = `
		SELECT
			id
			first_name
			last_name
			middle_name
			date_of_birth
			ish_joyi
			mutahassislik
			fan_tarmogi
			mavzular
			email
			password
			phone_number
			created_at
			updated_at
		FROM oqituvchi
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&first_name,
		&last_name,
		&middle_name,
		&date_of_birth,
		&ish_joyi,
		&mutahassislik,
		&fan_tarmogi,
		&mavzular,
		&email,
		&password,
		&phone_number,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Oqituvchi{
		Id:            id.String,
		FirstName:     first_name.String,
		LastName:      last_name.String,
		MiddleName:    middle_name.String,
		DateOfBirth:   date_of_birth.String,
		IshJoyi:       ish_joyi.String,
		Mutahassislik: mutahassislik.String,
		FanTarmogi:    fan_tarmogi.String,
		Mavzular:      mavzular.String,
		Email:         email.String,
		Password:      password.String,
		PhoneNumber:   phone_number.String,
		CreatedAt:     created_at.String,
		UpdatedAt:     updated_at.String,
	}, nil
}

func (r *oqituvchiRepo) GetList(ctx context.Context, req *models.OqituvchiGetListRequest) (*models.OqituvchiGetListResponse, error) {

	var (
		resp   = &models.OqituvchiGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			middle_name,
			date_of_birth,
			ish_joyi,
			mutahassislik,
			fan_tarmogi,
			mavzular,
			email,
			password,
			phone_number,
			created_at,
			updated_at
			
		FROM oqituvchi
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.SearchFirstName != "" {
		where += ` AND first_name ILIKE '%' || '` + req.SearchFirstName + `' || '%'`
	}
	if req.SearchLastName != "" {
		where += ` AND first_name ILIKE '%' || '` + req.SearchLastName + `' || '%'`
	}
	if req.SearchMiddleName != "" {
		where += ` AND first_name ILIKE '%' || '` + req.SearchMiddleName + `' || '%'`
	}
	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id            sql.NullString
			first_name    sql.NullString
			last_name     sql.NullString
			middle_name   sql.NullString
			date_of_birth sql.NullString
			ish_joyi      sql.NullString
			mutahassislik sql.NullString
			fan_tarmogi   sql.NullString
			mavzular      sql.NullString
			email         sql.NullString
			password      sql.NullString
			phone_number  sql.NullString
			created_at    sql.NullString
			updated_at    sql.NullString
		)

		err := rows.Scan(
			&id,
			&first_name,
			&last_name,
			&middle_name,
			&date_of_birth,
			&ish_joyi,
			&mutahassislik,
			&fan_tarmogi,
			&mavzular,
			&email,
			&password,
			&phone_number,
			&created_at,
			&updated_at,
		)

		if err != nil {
			return nil, err
		}

		resp.Oqituvchilar = append(resp.Oqituvchilar, &models.Oqituvchi{
			Id:            id.String,
			FirstName:     first_name.String,
			LastName:      last_name.String,
			MiddleName:    middle_name.String,
			DateOfBirth:   date_of_birth.String,
			IshJoyi:       ish_joyi.String,
			Mutahassislik: mutahassislik.String,
			FanTarmogi:    fan_tarmogi.String,
			Mavzular:      mavzular.String,
			Email:         email.String,
			Password:      password.String,
			PhoneNumber:   phone_number.String,
			CreatedAt:     created_at.String,
			UpdatedAt:     updated_at.String,
		})
	}

	return resp, nil
}

func (r *oqituvchiRepo) Update(ctx context.Context, req *models.UpdateOqituvchi) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			oqituvchi
		SET
		id = :id,
		first_name = :first_name,
		last_name = :last_name,
		middle_name = :middle_name,
		date_of_birth = :date_of_birth,
		ish_joyi = :ish_joyi,
		mutahassislik = :mutahassislik,
		fan_tarmogi = :fan_tarmogi,
		mavzular = :mavzular,
		email = :email,
		password = :password,
		phone_number = :phone_number,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"first_name":    req.FirstName,
		"last_name":     req.LastName,
		"middle_name":   req.MiddleName,
		"date_of_birth": req.DateOfBirth,
		"ish_joyi":      req.IshJoyi,
		"mutahassislik": req.Mutahassislik,
		"fan_tarmogi":   req.FanTarmogi,
		"mavzular":      req.Mavzular,
		"email":         req.Email,
		"password":      req.Password,
		"phone_number":  req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *oqituvchiRepo) Delete(ctx context.Context, req *models.OqituvchiPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM oqituvchi WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
