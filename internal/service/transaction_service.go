package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/gateway"
	"Payment-Gateway/internal/models"
	"Payment-Gateway/internal/repository"
	"time"

	"github.com/google/uuid"
)

type TransactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repository: repo,
	}
}

func (s *TransactionService) CreateAndProcessDeposit(req *models.DepositRequest, gw gateway.PaymentGateway) (*models.Transaction, error) {
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

	// Forward to gateway
	resp, err := gw.ProcessDeposit(nil)
	if err != nil {
		s.repository.UpdateTransactionStatus(tx.ID, constants.StatusFailed)
		return tx, err
	}

	s.repository.UpdateTransactionStatus(tx.ID, constants.StatusSuccess)
	_ = resp // In real code, process resp
	return tx, nil
}

func (s *TransactionService) CreateAndProcessWithdrawal(req *models.WithdrawalRequest, gw gateway.PaymentGateway) (*models.Transaction, error) {
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

	resp, err := gw.ProcessWithdrawal(nil)
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
