package models

type UserPrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	//Username string `json:"username"`
}

type CreateUser struct {
	FullName    string `json:"fullname"`
	Institution string `json:"institution"`
	Department  string `json:"department"`
	Degree      string `json:"degree"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	//Username string `json:"username"`

}

type User struct {
	Id          string `json:"id"`
	FullName    string `json:"fullname"`
	Institution string `json:"Institution"`
	Department  string `json:"department"`
	Degree      string `json:"degree"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateUser struct {
	Id          string `json:"id"`
	FullName    string `json:"fullname"`
	Institution string `json:"Institution"`
	Department  string `json:"department"`
	Degree      string `json:"degree"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	//Username string `json:"username"`

}

type UserGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type UserGetListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
