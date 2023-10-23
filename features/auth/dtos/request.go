package dtos

type InputUsers struct {
	Fullname    string `json:"fullname" form:"fullname"`
	PhoneNumber string `json:"phone_number" form:phone_number"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

type LoginUsers struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}