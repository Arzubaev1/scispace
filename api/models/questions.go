package models

type QuestionPrimaryKey struct {
	Id string `json:"id"`
}

type CreateQuestion struct {
	Title   string `json:"title"`
	Context string `json:"context"`
	Tags    string `json:"tags"`
}

type Question struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Context   string `json:"context"`
	Tags      string `json:"tags"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateQuestion struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Context string `json:"context"`
	Tags    string `json:"tags"`
}

type QuestionGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type QuestionGetListResponse struct {
	Count     int         `json:"count"`
	Questions []*Question `json:"questions"`
}
