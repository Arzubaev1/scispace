package models

type MaqolaPrimaryKey struct {
	Id string `json:"id"`
}
type Maqola struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Tavsifi          string `json:"tavsifi"`
	QoshimchaLinklar string `json:"qoshimcha_linklar"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
type CreateMaqola struct {
	Name             string `json:"name"`
	Tavsifi          string `json:"tavsifi"`
	QoshimchaLinklar string `json:"qoshimcha_linklar"`
}
type UpdateMaqola struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Tavsifi          string `json:"tavsifi"`
	QoshimchaLinklar string `json:"qoshimcha_linklar"`
}
type MaqolaGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type MaqolaGetListResponse struct {
	Count     int       `json:"count"`
	Maqolalar []*Maqola `json:"maqolalar"`
}
