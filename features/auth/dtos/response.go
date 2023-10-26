package dtos

import "shortlink/features/goly"

type ResUsers struct {
	ID          int    `json:"id`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Goly        []goly.Goly
}

type ResRegister struct {
	ID          int            `json:"id`
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
