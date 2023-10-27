package usecase

import (
	"errors"
	"shortlink/features/donate"
	"shortlink/features/donate/dtos"
	"shortlink/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
	"github.com/sirupsen/logrus"
)

type service struct {
	model donate.Repository
	validator *validator.Validate
}

func New(model donate.Repository, validate *validator.Validate) donate.Usecase {
	return &service{
		model: model,
		validator: validate,
	}
}

func (svc *service) FindAll(page, size int) []dtos.TransactionInputResponse{
	var donate []dtos.TransactionInputResponse

	donateEnt := svc.model.Paginate(page, size)

	for _, donated := range donateEnt {
		var data dtos.TransactionInputResponse

		if err := smapping.FillStruct(&data, smapping.MapFields(donated)); err != nil {
			log.Error(err.Error())
		}

		donate = append(donate, data)
	}

	return donate
}

func (svc *service) Create(newData dtos.TransactionInput) (*dtos.TransactionInputResponse, error) {
	
	err := svc.validator.Struct(newData)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var newTransaction = helpers.RequestToTransaction(newData)
	newTransaction.OrderID = helpers.GenerateUUID()
	result := svc.model.Insert(newTransaction)
	if result == nil {
        return nil, errors.New("failed to insert transaction")
    }
	
	token, url := svc.model.SnapRequest(result.OrderID, int64(newData.Amount))
	var TransactionInputResponse = helpers.TransactionToResponseInput(result, token, url)

	return TransactionInputResponse, nil
}

func (svc *service) Notifications(notificationPayload map[string]any) error {
	orderID, exist := notificationPayload["order_id"].(string)

	if !exist {
		logrus.Error("order id not found in notification payload")
		return errors.New("invalid notification payload")
	}

	status, err := svc.model.CheckTransaction(orderID)
	if err != nil {
		logrus.Errorf("Error checking transaction for OrderID %s: %v", orderID, err)
		return err
	}

	transaction, _ := svc.model.SelectByOrderID(orderID)

	err = svc.model.UpdateStatusTransaction(transaction.ID, status.Transaction)
	if err != nil {
		logrus.Errorf("Error updating order status for OrderID %s: %v", orderID, err)
		return err
	}
	return nil
}