package helpers

import (
	"shortlink/features/donate"
	"shortlink/features/donate/dtos"
)

func RequestToTransaction(data dtos.TransactionInput) *donate.Transaction {
	return &donate.Transaction{
		ID: data.OrderID,
		UserID: data.UserID,
		Amount: int64(data.Amount),
		Description: data.Description,
	}
}

func TransactionToResponseInput(data *donate.Transaction, token string, url string) *dtos.TransactionInputResponse {
	return &dtos.TransactionInputResponse{
		ID:          data.ID,
		UserID: data.UserID,
		OrderID:     data.OrderID,
		Status:      data.Status,
		Amount: uint(data.Amount),
		Description: data.Description,
		Token:       token,
		RedirectURL: url,
	}
}

func TransactionToResponse(data *donate.Transaction) *dtos.TransactionResponse {
	return &dtos.TransactionResponse{
		ID:      data.ID,
		OrderID: data.OrderID,
		Status:  data.Status,
	}
}