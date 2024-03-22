package midtrans

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService struct {
	client snap.Client
}

func NewMidtrans() MidtransServiceInterface {
	var client snap.Client
	client.New("SB-Mid-server-JdUsS8ZWJTenkX9It0AQRGz-", midtrans.Sandbox)
	return &MidtransService{
		client: client,
	}
}

func (m MidtransService) GenerateSnapURL(ctx context.Context, totalamount int64) (string, string, error) {
	// 1. Initiate Snap client
	uuid := uuid.New()
	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.String(),
			GrossAmt: totalamount,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, err := m.client.CreateTransaction(req)
	if err != nil {
		return "", "", err
	}
	fmt.Println("Response :")

	return snapResp.RedirectURL, req.TransactionDetails.OrderID, nil
}
func (m MidtransService) VerifyPayment(ctx context.Context, order_id string) (string, error) {
	var client coreapi.Client
	client.New("SB-Mid-server-JdUsS8ZWJTenkX9It0AQRGz-", midtrans.Sandbox)
	// 1. Initialize empty map

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(order_id)
	if e != nil {
		return "", e
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
					return "success", nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				return "success", nil
				// TODO set transaction status on your databaase to 'success'
			} else if transactionStatusResp.TransactionStatus == "deny" {
				return "deny", nil
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
				return "failure", nil
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
				return "pending", nil
			}
		}
	}
	return "", nil
}
