package models

type PostPrimaryKey struct {
	Id string `json:"id"`
}

type CreatePost struct {
	Title   string `json:"title"`
	Context string `json:"context"`
	Link    string `json:"link"`
}

type Post struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Context   string `json:"context"`
	Link      string `json:"link"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdatePost struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Context string `json:"context"`
	Link    string `json:"link"`
}

type PostGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type PostGetListResponse struct {
	Count int     `json:"count"`
	Posts []*Post `json:"posts"`
}
