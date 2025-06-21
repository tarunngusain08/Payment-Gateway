package models

import (
	"Payment-Gateway/internal/constants"
	"time"
)

type Transaction struct {
	ID        string                      `json:"id"`
	Type      constants.TransactionType   `json:"type"`
	Amount    float64                     `json:"amount"`
	Status    constants.TransactionStatus `json:"status"`
	Timestamp time.Time                   `json:"timestamp"`
	Account   string                      `json:"account"`
}

type DepositRequest struct {
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

type WithdrawalRequest struct {
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}
