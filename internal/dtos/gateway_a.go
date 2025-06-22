package dtos

import errors "Payment-Gateway/pkg/error"

type GatewayADepositRequest struct {
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

func (r *GatewayADepositRequest) Validate() error {
	if r.Account == "" {
		return errors.ErrAccountRequired
	}
	if r.Amount <= 0 {
		return errors.ErrAmountMustBePositive
	}
	return nil
}

type GatewayAWithdrawalRequest struct {
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

func (r *GatewayAWithdrawalRequest) Validate() error {
	if r.Account == "" {
		return errors.ErrAccountRequired
	}
	if r.Amount <= 0 {
		return errors.ErrAmountMustBePositive
	}
	return nil
}
