package dtos

type GolyResponse struct {
	ID       uint64 `json:"id"`
	UserID   string `json:"user_id"`
	Redirect string `json:"redirect"`
	Goly     string `json:"goly"`
	Clicked  uint64 `json:"clicked"`
}