package dtos

type TransactionResponse struct {
	ID      uint   `json:"id"`
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
	Amount  uint   `json:"amount"`
}

type TransactionInputResponse struct {
	ID          uint   `json:"id"`
	OrderID     string `json:"order_id"`
	Status      string `json:"status"`
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}