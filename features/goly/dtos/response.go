package dtos

import "time"

type GolyResponse struct {
	ID         uint64   `json:"id"`
	UserID     string   `json:"user_id"`
	Redirect   string   `json:"redirect"`
	Custom     string   `json:"custom"`
	Clicked    uint64   `json:"clicked"`
	ExpiryDate time.Time `json:"expiry_date"`
}

type IPAddressResponse struct {
	Goly_id string `json:"goly_id"`
	IP      string `json:"ip"`
}