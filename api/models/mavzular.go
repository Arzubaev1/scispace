package models

type MavzuPrimaryKey struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type Mavzu struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type CreateMavzu struct {
	Name string `json:"name"`
}
type UpdateMavzu struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type MavzuGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type MavzuGetListResponse struct {
	Count    int      `json:"count"`
	Mavzular []*Mavzu `json:"mavzular"`
}
