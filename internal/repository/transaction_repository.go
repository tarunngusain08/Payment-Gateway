package repository

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/models"
	"sync"
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
	return &InMemoryTransactionRepository{}
}

func (r *InMemoryTransactionRepository) CreateTransaction(tx *models.Transaction) error {
	r.store.Store(tx.ID, tx)
	return nil
}

func (r *InMemoryTransactionRepository) UpdateTransactionStatus(id string, status constants.TransactionStatus) error {
	val, ok := r.store.Load(id)
	if !ok {
		return nil
	}
	tx := val.(*models.Transaction)
	tx.Status = status
	r.store.Store(id, tx)
	return nil
}

func (r *InMemoryTransactionRepository) GetTransactionByID(id string) (*models.Transaction, bool) {
	val, ok := r.store.Load(id)
	if !ok {
		return nil, false
	}
	return val.(*models.Transaction), true
}
