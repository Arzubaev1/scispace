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

type otherRepo struct {
	db *pgxpool.Pool
}

func NewOtherRepo(db *pgxpool.Pool) *otherRepo {
	return &otherRepo{
		db: db,
	}
}

// FullName    string `json:"fullname"`
// 	Institution string `json:"institution"`
// 	Department  string `json:"department"`
// 	Degree      string `json:"degree"`
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`

func (r *otherRepo) Create(ctx context.Context, req *models.CreateOther) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO other(
			id,
			first_name,
			last_name,
			middle_name,
			date_of_birth,
			oqish_joyi,
			yonalish,
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
		req.OqishJoyi,
		req.Yonalish,
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

func (r *otherRepo) GetByID(ctx context.Context, req *models.OtherPrimaryKey) (*models.Other, error) {

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
		oqish_joyi    sql.NullString
		yonalish      sql.NullString
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
			oqish_joyi
			yonalish
			mavzular
			email
			password
			phone_number
			created_at
			updated_at
		FROM other
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&first_name,
		&last_name,
		&middle_name,
		&date_of_birth,
		&oqish_joyi,
		&yonalish,
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

	return &models.Other{
		Id:          id.String,
		FirstName:   first_name.String,
		LastName:    last_name.String,
		MiddleName:  middle_name.String,
		DateOfBirth: date_of_birth.String,
		OqishJoyi:   oqish_joyi.String,
		Yonalish:    yonalish.String,
		Mavzular:    mavzular.String,
		Email:       email.String,
		Password:    password.String,
		PhoneNumber: phone_number.String,
		CreatedAt:   created_at.String,
		UpdatedAt:   updated_at.String,
	}, nil
}

func (r *otherRepo) GetList(ctx context.Context, req *models.OtherGetListRequest) (*models.OtherGetListResponse, error) {

	var (
		resp   = &models.OtherGetListResponse{}
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
			oqish_joyi,
			yonalish,
			mavzular,
			email,
			password,
			phone_number,
			created_at,
			updated_at
			
		FROM other
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
			oqish_joyi    sql.NullString
			yonalish      sql.NullString
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
			&oqish_joyi,
			&yonalish,
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

		resp.Others = append(resp.Others, &models.Other{
			Id:          id.String,
			FirstName:   first_name.String,
			LastName:    last_name.String,
			MiddleName:  middle_name.String,
			DateOfBirth: date_of_birth.String,
			OqishJoyi:   oqish_joyi.String,
			Yonalish:    yonalish.String,
			Mavzular:    mavzular.String,
			Email:       email.String,
			Password:    password.String,
			PhoneNumber: phone_number.String,
			CreatedAt:   created_at.String,
			UpdatedAt:   updated_at.String,
		})
	}

	return resp, nil
}

func (r *otherRepo) Update(ctx context.Context, req *models.UpdateOther) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			other
		SET
		id = :id,
		first_name = :first_name,
		last_name = :last_name,
		middle_name = :middle_name,
		date_of_birth = :date_of_birth,
		oqish_joyi = :oqish_joyi,
		yonalish = :fan_tarmogi,
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
		"oqish_joyi":    req.OqishJoyi,
		"yonalish":      req.Yonalish,
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

func (r *otherRepo) Delete(ctx context.Context, req *models.OtherPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM other WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
