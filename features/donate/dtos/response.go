package dtos

type TransactionResponse struct {
	ID          uint   `json:"id"`
	UserID      string `json:"user_id"`
	OrderID     string `json:"order_id"`
	Status      string `json:"status"`
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
}

type TransactionInputResponse struct {
	ID          uint   `json:"id"`
	UserID      string `json:"user_id"`
	OrderID     string `json:"order_id"`
	Status      string `json:"status"`
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}