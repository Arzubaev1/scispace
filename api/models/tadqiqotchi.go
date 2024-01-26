package models

type TadqiqotchiPrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
type Tadqiqotchi struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`
	OqishJoyi   string `json:"oqish_joyi"`
	FanTarmogi  string `json:"fan_tarmogi"`
	Mavzular    string `json:"mavzular"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type CreateTadqiqotchi struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`
	OqishJoyi   string `json:"oqish_joyi"`
	FanTarmogi  string `json:"FanTarmogi"`
	Mavzular    string `json:"mavzular"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
type UpdateTadqiqotchi struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`
	OqishJoyi   string `json:"oqish_joyi"`
	FanTarmogi  string `json:"FanTarmogi"`
	Mavzular    string `json:"mavzular"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type TadqiqotchiGetListRequest struct {
	Offset           int    `json:"offset"`
	Limit            int    `json:"limit"`
	SearchFirstName  string `json:"search_first_name"`
	SearchLastName   string `json:"search_last_name"`
	SearchMiddleName string `json:"search_middle_name"`
}

type TadqiqotchiGetListResponse struct {
	Count          int            `json:"count"`
	Tadqiqotchilar []*Tadqiqotchi `json:"tadqiqotchilar"`
}
