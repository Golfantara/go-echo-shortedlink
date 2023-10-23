package dtos

type ResUsers struct {
	ID          int    `json:"id`
	UsersID     string `json:"user_id"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type ResRegister struct {
	ID          int            `json:"id`
	UsersID     string         `json:"users_id"`
	Fullname    string         `json:"fullname"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	Token       map[string]any `json:"token"`
}

type ResLogin struct {
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Token map[string]any `json:"token"`
}
