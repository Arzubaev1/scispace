package models

type MutahassislikPrimaryKey struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type Mutahassislik struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type CreateMutahassislik struct {
	Name string `json:"name"`
}
type UpdateMutahassislik struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type MutahassislikGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type MutahassislikGetListResponse struct {
	Count            int              `json:"count"`
	Mutahassisliklar []*Mutahassislik `json:"mutahassisliklar"`
}
