package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/gateway"
	"Payment-Gateway/internal/models"
)

type Callback interface {
	HandleCallback(req dtos.HandleCallbackRequest) (*dtos.HandleCallbackResponse, error)
}

type Deposit interface {
	CreateAndProcessDeposit(req *models.DepositRequest) (*models.Transaction, error)
}

type Withdrawal interface {
	CreateAndProcessWithdrawal(req *models.WithdrawalRequest) (*models.Transaction, error)
}

type GatewayPool interface {
	GetAllGateways() ([]gateway.PaymentGateway, error)
	GetRoundRobinGateway() (gateway.PaymentGateway, error)
}

type Transaction interface {
	UpdateStatus(id string, status constants.TransactionStatus) error
	Deposit
	Withdrawal
}
