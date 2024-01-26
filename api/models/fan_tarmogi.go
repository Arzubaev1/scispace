package models

type FanTarmogiPrimaryKey struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type FanTarmogi struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type CreateFanTarmogi struct {
	Name string `json:"name"`
}
type UpdateFanTarmogi struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type FanTarmogiGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type FanTarmogiGetListResponse struct {
	Count         int           `json:"count"`
	FanTarmoqlari []*FanTarmogi `json:"FanTarmoqlari"`
}
