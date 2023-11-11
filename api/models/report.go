package models

type ReportPrimaryKey struct {
	Id string `json:"id"`
}

type CreateReport struct {
	ReportStatus string `json:"report_status"`
	Description  string `json:"description"`
}

type Report struct {
	Id           string `json:"id"`
	ReportStatus string `json:"report_status"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UpdateReport struct {
	Id           string `json:"id"`
	ReportStatus string `json:"report_status"`
	Description  string `json:"description"`
}

type ReportGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ReportGetListResponse struct {
	Count   int       `json:"count"`
	Reports []*Report `json:"reports"`
}
