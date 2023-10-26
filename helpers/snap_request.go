package helpers

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateSnapRequest(snapClient snap.Client, orderID string, amount int64) (*snap.Response, error) {
	var snapRequest = &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
            OrderID:    orderID,
            GrossAmt: amount,
        },
		EnabledPayments: snap.AllSnapPaymentType,
	}
	snapResponse, err := snapClient.CreateTransaction(snapRequest)
	if err != nil {
		return nil, err
	}
	return snapResponse, nil
}