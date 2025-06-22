package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/models"
	"Payment-Gateway/internal/repository"
	"Payment-Gateway/pkg/logger"
	"time"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type TransactionService struct {
	repository repository.TransactionRepository
	Gateway    GatewayPool
}

func NewTransactionService(repo repository.TransactionRepository, gateway GatewayPool) Transaction {
	return &TransactionService{
		repository: repo,
		Gateway:    gateway,
	}
}

func (s *TransactionService) CreateAndProcessDeposit(req *models.DepositRequest) (*models.Transaction, error) {
	log := logger.GetLogger().With(
		zap.String("func", "TransactionService.CreateAndProcessDeposit"),
		zap.String("account", req.Account),
		zap.Float64("amount", req.Amount),
	)

	log.Info("Creating deposit transaction")
	tx := &models.Transaction{
		ID:        uuid.NewString(),
		Type:      constants.TypeDeposit,
		Amount:    req.Amount,
		Status:    constants.StatusPending,
		Timestamp: time.Now(),
		Account:   req.Account,
	}

	if err := s.repository.CreateTransaction(tx); err != nil {
		log.Error("Failed to create transaction", zap.Error(err))
		return nil, err
	}

	gateway, err := s.Gateway.GetRoundRobinGateway()
	if err != nil {
		log.Error("No gateway available", zap.Error(err))
		return nil, err
	}

	log.Info("Processing deposit with gateway")
	resp, err := gateway.ProcessDeposit(nil)
	if err != nil {
		log.Error("Gateway deposit failed", zap.Error(err))
		s.repository.UpdateTransactionStatus(tx.ID, constants.StatusFailed)
		return tx, err
	}

	s.repository.UpdateTransactionStatus(tx.ID, constants.StatusSuccess)
	log.Info("Deposit processed successfully", zap.Any("gateway_response", resp))
	return tx, nil
}

func (s *TransactionService) CreateAndProcessWithdrawal(req *models.WithdrawalRequest) (*models.Transaction, error) {
	log := logger.GetLogger().With(
		zap.String("func", "TransactionService.CreateAndProcessWithdrawal"),
		zap.String("account", req.Account),
		zap.Float64("amount", req.Amount),
	)

	log.Info("Creating withdrawal transaction")
	tx := &models.Transaction{
		ID:        uuid.NewString(),
		Type:      constants.TypeWithdrawal,
		Amount:    req.Amount,
		Status:    constants.StatusPending,
		Timestamp: time.Now(),
		Account:   req.Account,
	}

	if err := s.repository.CreateTransaction(tx); err != nil {
		log.Error("Failed to create transaction", zap.Error(err))
		return nil, err
	}

	gateway, err := s.Gateway.GetRoundRobinGateway()
	if err != nil {
		log.Error("No gateway available", zap.Error(err))
		return nil, err
	}

	log.Info("Processing withdrawal with gateway")
	resp, err := gateway.ProcessWithdrawal(nil)
	if err != nil {
		log.Error("Gateway withdrawal failed", zap.Error(err))
		s.repository.UpdateTransactionStatus(tx.ID, constants.StatusFailed)
		return tx, err
	}

	err = s.repository.UpdateTransactionStatus(tx.ID, constants.StatusSuccess)
	if err != nil {
		log.Error("Failed to update transaction status", zap.Error(err))
		return tx, err
	}
	log.Info("Withdrawal processed successfully", zap.Any("gateway_response", resp))
	return tx, nil
}

func (s *TransactionService) GetTransaction(id string) (*models.Transaction, bool) {
	log := logger.GetLogger().With(
		zap.String("func", "TransactionService.GetTransaction"),
		zap.String("transaction_id", id),
	)
	tx, found := s.repository.GetTransactionByID(id)
	if found {
		log.Info("Transaction found")
	} else {
		log.Warn("Transaction not found")
	}
	return tx, found
}

// Implement UpdateStatus method
func (s *TransactionService) UpdateStatus(id string, status constants.TransactionStatus) error {
	log := logger.GetLogger().With(
		zap.String("func", "TransactionService.UpdateStatus"),
		zap.String("transaction_id", id),
		zap.String("status", string(status)),
	)
	log.Info("Updating transaction status")
	err := s.repository.UpdateTransactionStatus(id, status)
	if err != nil {
		log.Error("Failed to update transaction status", zap.Error(err))
	}
	return err
}
