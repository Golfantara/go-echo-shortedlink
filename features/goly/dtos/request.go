package dtos

type CreateGolyInput struct {
	UsersID  string `json:"users_id"`
	Goly     string `json:"goly" form:"goly"`
	Redirect string `json:"redirect" form:"redirect"`
	Random   bool   `json:"random" form:"random"`
}