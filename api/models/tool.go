package models

type ToolPrimaryKey struct {
	Id string `json:"id"`
}

type CreateTool struct {
	ToolName string `json:"tool_name"`
	Link     string `json:"link"`
}

type Tool struct {
	Id        string `json:"id"`
	ToolName  string `json:"tool_name"`
	Link      string `json:"link"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateTool struct {
	Id       string `json:"id"`
	ToolName string `json:"tool_name"`
	Link     string `json:"link"`
}

type ToolGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ToolGetListResponse struct {
	Count int     `json:"count"`
	Tools []*Tool `json:"tools"`
}
