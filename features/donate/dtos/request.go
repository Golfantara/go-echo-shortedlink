package dtos

type TransactionInput struct {
	OrderID uint `json:"order_id" form:"order_id"`
	Amount  uint `json:"amount" form:"amount"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}