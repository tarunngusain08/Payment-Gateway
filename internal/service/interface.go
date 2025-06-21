package service

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/gateway"
	"Payment-Gateway/internal/models"
)

type Callback interface {
	HandleCallback(req dtos.HandleCallbackRequest) (*dtos.HandleCallbackResponse, error)
}

type Deposit interface {
	CreateAndProcessDeposit(req *models.DepositRequest, gw gateway.PaymentGateway) (*models.Transaction, error)
}

type Withdrawal interface {
	CreateAndProcessWithdrawal(req *models.WithdrawalRequest, gw gateway.PaymentGateway) (*models.Transaction, error)
}
