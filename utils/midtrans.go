package utils

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransSnapClient(serverKey string) snap.Client {
	var snapClient snap.Client
	snapClient.New(serverKey, midtrans.Sandbox)

	return snapClient
}

func MidtransCoreAPIClient(serverKey string) coreapi.Client {
	var coreAPIClient coreapi.Client
	coreAPIClient.New(serverKey, midtrans.Sandbox)

	return coreAPIClient
}