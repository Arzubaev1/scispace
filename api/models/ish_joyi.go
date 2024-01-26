package models

type IshJoyiPrimaryKey struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type IshJoyi struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type CreateIshJoyi struct {
	Name string `json:"name"`
}
type UpdateIshJoyi struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type IshJoyiGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type IshJoyiGetListResponse struct {
	Count      int        `json:"count"`
	IshJoylari []*IshJoyi `json:"ish_joylari"`
}
