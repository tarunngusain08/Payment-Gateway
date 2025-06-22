package gateway

import (
	"net/http"
)

// PaymentGateway defines the contract for all payment gateway integrations.
type PaymentGateway interface {
	ProcessDeposit(r *http.Request) (interface{}, error)
	ProcessWithdrawal(r *http.Request) (interface{}, error)
}
