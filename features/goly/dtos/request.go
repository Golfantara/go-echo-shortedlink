package dtos

type CreateGolyInput struct {
	UserID       string `json:"user_id" form:"user_id"`
	Custom       string `json:"custom" form:"custom"`
	Redirect     string `json:"redirect" form:"redirect"`
	Random       bool   `json:"random" form:"random"`
	ExpiryInDays int    `json:"expiry_in_days" form:"expiry_in_days"`
}
type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}