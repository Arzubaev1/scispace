package models

type DatabasePrimaryKey struct {
	Id string `json:"id"`
}

type CreateDatabase struct {
	DatabaseName string `json:"database_name"`
	Link         string `json:"link"`
}

type Database struct {
	Id           string `json:"id"`
	DatabaseName string `json:"database_name"`
	Link         string `json:"link"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UpdateDatabase struct {
	Id           string `json:"id"`
	DatabaseName string `json:"database_name"`
	Link         string `json:"link"`
}

type DatabaseGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type DatabaseGetListResponse struct {
	Count     int         `json:"count"`
	Databases []*Database `json:"databases"`
}
