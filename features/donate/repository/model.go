package repository

import (
	"errors"
	"shortlink/features/donate"
	"shortlink/helpers"

	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
	snapClient snap.Client
	coreAPIClient coreapi.Client
}

func New(db *gorm.DB, snapClient snap.Client, coreAPIClient coreapi.Client) donate.Repository {
	return &model{
        db: db,
		snapClient: snapClient,
		coreAPIClient: coreAPIClient,
    }
}

func (mdl *model) Paginate(page, size int) []donate.Transaction {
	var transaction []donate.Transaction

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&transaction)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return transaction
}

func (mdl *model) Insert(newData *donate.Transaction) *donate.Transaction {
	result := mdl.db.Create(newData)
	

	if result.Error!= nil {
        logrus.Errorf("Error inserting transaction: %v", result.Error)
		return nil
    }
	return newData
}

func (mdl *model) SelectByID(id int) *donate.Transaction {
	var transaction donate.Transaction
	result := mdl.db.Where("id = ?", id).Find(&transaction)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &transaction
}

func (mdl *model) SelectByOrderID(orderID string) (*donate.Transaction, error) {
	var transaction donate.Transaction
	result := mdl.db.Table("transactions").Where("order_id = ?", orderID).Find(&transaction)

	if result.Error != nil {
		logrus.Error("Respository : Get transaction by id error,", result.Error)
		return nil, result.Error
	}
	return &transaction, nil
}

func (mdl *model) DeleteByID(userID int) int64 {
	result := mdl.db.Delete(&donate.Transaction{}, userID)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}

func (mdl *model) SnapRequest(orderID string, amount int64) (string, string) {
	snapResponse, err := helpers.CreateSnapRequest(mdl.snapClient, orderID, amount)
	if err != nil {
		return "", ""
	}

	return snapResponse.Token, snapResponse.RedirectURL
}

func (mdl *model) CheckTransaction(orderID string) (donate.Status, error) {
	var status donate.Status

	transactionStatusResp, err := mdl.coreAPIClient.CheckTransaction(orderID)
	if err != nil {
		return donate.Status{}, err
	} else {
		if transactionStatusResp != nil {
			status = helpers.TransactionStatus(transactionStatusResp)
			return status, nil
		}
	}
	return donate.Status{}, err
}

func (mdl *model) UpdateStatusTransaction(id uint, status string) error {
	result := mdl.db.Table("transactions").Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		logrus.Error("Repository: Update transaction status error,", result.Error)
		return result.Error
	}

	if result.RowsAffected < 1 {
		logrus.Error("Repository: No row Affected ,", result.Error)
		return errors.New("data not found")
	}

	return nil
}