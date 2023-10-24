package dtos

type CreateGolyInput struct {
	UserID   string `json:"user_id" form:"user_id"`
	Goly     string `json:"goly" form:"goly"`
	Redirect string `json:"redirect" form:"redirect"`
	Random   bool   `json:"random" form:"random"`
}
type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}