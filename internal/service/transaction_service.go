package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/models"
	"Payment-Gateway/internal/repository"
	"time"

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
	tx := &models.Transaction{
		ID:        uuid.NewString(),
		Type:      constants.TypeDeposit,
		Amount:    req.Amount,
		Status:    constants.StatusPending,
		Timestamp: time.Now(),
		Account:   req.Account,
	}

	if err := s.repository.CreateTransaction(tx); err != nil {
		return nil, err
	}

	gateway, err := s.Gateway.GetRoundRobinGateway()
	if err != nil {
		return nil, err
	}

	// Forward to gateway
	resp, err := gateway.ProcessDeposit(nil)
	if err != nil {
		s.repository.UpdateTransactionStatus(tx.ID, constants.StatusFailed)
		return tx, err
	}

	s.repository.UpdateTransactionStatus(tx.ID, constants.StatusSuccess)
	_ = resp // In real code, process resp
	return tx, nil
}

func (s *TransactionService) CreateAndProcessWithdrawal(req *models.WithdrawalRequest) (*models.Transaction, error) {
	tx := &models.Transaction{
		ID:        uuid.NewString(),
		Type:      constants.TypeWithdrawal,
		Amount:    req.Amount,
		Status:    constants.StatusPending,
		Timestamp: time.Now(),
		Account:   req.Account,
	}

	if err := s.repository.CreateTransaction(tx); err != nil {
		return nil, err
	}

	gateway, err := s.Gateway.GetRoundRobinGateway()
	if err != nil {
		return nil, err
	}

	resp, err := gateway.ProcessWithdrawal(nil)
	if err != nil {
		s.repository.UpdateTransactionStatus(tx.ID, constants.StatusFailed)
		return tx, err
	}

	err = s.repository.UpdateTransactionStatus(tx.ID, constants.StatusSuccess)
	if err != nil {
		return tx, err
	}
	_ = resp
	return tx, nil
}

func (s *TransactionService) GetTransaction(id string) (*models.Transaction, bool) {
	return s.repository.GetTransactionByID(id)
}

// Implement UpdateStatus method
func (s *TransactionService) UpdateStatus(id string, status constants.TransactionStatus) error {
	return s.repository.UpdateTransactionStatus(id, status)
}
