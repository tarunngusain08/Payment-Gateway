package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/gateway"
	"Payment-Gateway/internal/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

type TransactionService struct {
	store map[string]*models.Transaction
	mu    sync.Mutex
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		store: make(map[string]*models.Transaction),
	}
}

// CreateAndProcessDeposit creates a PENDING transaction, forwards to gateway, updates status.
func (s *TransactionService) CreateAndProcessDeposit(req *models.DepositRequest, gw gateway.PaymentGateway) (*models.Transaction, error) {
	tx := &models.Transaction{
		ID:        uuid.NewString(),
		Type:      constants.TypeDeposit,
		Amount:    req.Amount,
		Status:    constants.StatusPending,
		Timestamp: time.Now(),
		Account:   req.Account,
	}
	s.mu.Lock()
	s.store[tx.ID] = tx
	s.mu.Unlock()

	// Forward to gateway
	// Simulate request (in real code, build http.Request)
	resp, err := gw.ProcessDeposit(nil)
	s.mu.Lock()
	defer s.mu.Unlock()
	if err != nil {
		tx.Status = constants.StatusFailed
		return tx, err
	}
	tx.Status = constants.StatusSuccess
	_ = resp // In real code, process resp
	return tx, nil
}

// CreateAndProcessWithdrawal creates a PENDING transaction, forwards to gateway, updates status.
func (s *TransactionService) CreateAndProcessWithdrawal(req *models.WithdrawalRequest, gw gateway.PaymentGateway) (*models.Transaction, error) {
	tx := &models.Transaction{
		ID:        uuid.NewString(),
		Type:      constants.TypeWithdrawal,
		Amount:    req.Amount,
		Status:    constants.StatusPending,
		Timestamp: time.Now(),
		Account:   req.Account,
	}
	s.mu.Lock()
	s.store[tx.ID] = tx
	s.mu.Unlock()

	// Forward to gateway
	resp, err := gw.ProcessWithdrawal(nil)
	s.mu.Lock()
	defer s.mu.Unlock()
	if err != nil {
		tx.Status = constants.StatusFailed
		return tx, err
	}
	tx.Status = constants.StatusSuccess
	_ = resp
	return tx, nil
}

// GetTransaction returns a transaction by ID.
func (s *TransactionService) GetTransaction(id string) (*models.Transaction, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	tx, ok := s.store[id]
	return tx, ok
}
