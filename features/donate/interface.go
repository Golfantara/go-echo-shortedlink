package donate

import (
	"shortlink/features/donate/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Transaction
	Insert(newData *Transaction) *Transaction
	SelectByID(userID int) *Transaction
	SelectByOrderID(orderID string) (*Transaction, error)
	DeleteByID(userID int) int64
	SnapRequest(orderID string, amount int64) (string, string)
	CheckTransaction(orderID string) (Status, error)
	UpdateStatusTransaction(id uint, status string) error
}

type Usecase interface {
	FindAll(page, size int) []dtos.TransactionInputResponse
	Create(newData dtos.TransactionInput) (*dtos.TransactionInputResponse, error)
	Notifications(notificationPayload map[string]any) error
}

type Handler interface {
	Insert() echo.HandlerFunc
	Notifications() echo.HandlerFunc
	GetAllDonated() echo.HandlerFunc
}