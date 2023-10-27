package helpers

import (
	"shortlink/features/donate"

	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/sirupsen/logrus"
)

func TransactionStatus(transactionStatusResp *coreapi.TransactionStatusResponse) donate.Status {
	var status donate.Status

	if transactionStatusResp.TransactionStatus == "capture" {
		if transactionStatusResp.FraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			status.Transaction = "challenge"
			status.Donate = "Unpaid"
		} else if transactionStatusResp.FraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			status.Transaction = "success"
			status.Donate = "Paid"
		}
	} else if transactionStatusResp.TransactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		status.Transaction = "success"
		status.Donate = "Paid"
	} else if transactionStatusResp.TransactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
	} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		status.Transaction = "failed"
		status.Donate = "failed"
	} else if transactionStatusResp.TransactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		status.Transaction = "pending"
		status.Donate = "unpaid"
	}

	logrus.Infof("TransactionStatus: For OrderID %s, Transaction Status: %s, Donate Status: %s", transactionStatusResp.OrderID, status.Transaction, status.Donate)
	return status
}