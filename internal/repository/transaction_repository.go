package repository

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/models"
	"Payment-Gateway/pkg/logger"
	"sync"

	"go.uber.org/zap"
)

type TransactionRepository interface {
	CreateTransaction(tx *models.Transaction) error
	UpdateTransactionStatus(id string, status constants.TransactionStatus) error
	GetTransactionByID(id string) (*models.Transaction, bool)
}

type InMemoryTransactionRepository struct {
	store sync.Map // map[string]*models.Transaction
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	log := logger.GetLogger().With(zap.String("func", "NewInMemoryTransactionRepository"))
	log.Info("Initializing in-memory transaction repository")
	return &InMemoryTransactionRepository{}
}

func (r *InMemoryTransactionRepository) CreateTransaction(tx *models.Transaction) error {
	log := logger.GetLogger().With(
		zap.String("func", "InMemoryTransactionRepository.CreateTransaction"),
		zap.String("transaction_id", tx.ID),
	)
	log.Info("Creating transaction")
	r.store.Store(tx.ID, tx)
	return nil
}

func (r *InMemoryTransactionRepository) UpdateTransactionStatus(id string, status constants.TransactionStatus) error {
	log := logger.GetLogger().With(
		zap.String("func", "InMemoryTransactionRepository.UpdateTransactionStatus"),
		zap.String("transaction_id", id),
		zap.String("status", string(status)),
	)
	val, ok := r.store.Load(id)
	if !ok {
		log.Warn("Transaction not found for update")
		return nil
	}
	tx := val.(*models.Transaction)
	tx.Status = status
	r.store.Store(id, tx)
	log.Info("Transaction status updated")
	return nil
}

func (r *InMemoryTransactionRepository) GetTransactionByID(id string) (*models.Transaction, bool) {
	log := logger.GetLogger().With(
		zap.String("func", "InMemoryTransactionRepository.GetTransactionByID"),
		zap.String("transaction_id", id),
	)
	val, ok := r.store.Load(id)
	if !ok {
		log.Warn("Transaction not found")
		return nil, false
	}
	log.Info("Transaction found")
	return val.(*models.Transaction), true
}
