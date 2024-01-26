package models

type OtherPrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
type Other struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`
	OqishJoyi   string `json:"oqish_joyi"`
	Yonalish    string `json:"yonalish"`
	Mavzular    string `json:"mavzular"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type CreateOther struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`
	OqishJoyi   string `json:"oqish_joyi"`
	Yonalish    string `json:"yonalish"`
	Mavzular    string `json:"mavzular"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
type UpdateOther struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`
	OqishJoyi   string `json:"oqish_joyi"`
	Yonalish    string `json:"Yonalish"`
	Mavzular    string `json:"mavzular"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type OtherGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	SearchFirstName  string `json:"search_first_name"`
	SearchLastName   string `json:"search_last_name"`
	SearchMiddleName string `json:"search_middle_name"`
}

type OtherGetListResponse struct {
	Count  int      `json:"count"`
	Others []*Other `json:"Others"`
}
