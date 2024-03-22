package midtrans

import (
	"context"
)

type MidtransServiceInterface interface {
	GenerateSnapURL(ctx context.Context, totalamount int64) (string, string, error)
	VerifyPayment(ctx context.Context, order_id string) (string, error)
}
