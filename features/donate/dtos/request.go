package dtos

type TransactionInput struct {
	UserID      string `json:"user_id" form:"user_id"`
	OrderID     uint   `json:"order_id" form:"order_id"`
	Amount      uint   `json:"amount" form:"amount"`
	Description string `json:"description" form:"description"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}