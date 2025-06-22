package repository

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/models"
	"testing"
)

func TestCreateAndGetTransaction(t *testing.T) {
	repo := NewInMemoryTransactionRepository()
	tx := &models.Transaction{
		ID:      "tx1",
		Type:    constants.TypeDeposit,
		Amount:  100,
		Status:  constants.StatusPending,
		Account: "acc1",
	}

	// Create transaction
	err := repo.CreateTransaction(tx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Get transaction
	got, ok := repo.GetTransactionByID("tx1")
	if !ok {
		t.Fatalf("expected transaction to exist")
	}
	if got.ID != tx.ID || got.Amount != tx.Amount || got.Status != tx.Status {
		t.Errorf("got %+v, want %+v", got, tx)
	}
}

func TestGetTransactionByID_NotFound(t *testing.T) {
	repo := NewInMemoryTransactionRepository()
	_, ok := repo.GetTransactionByID("not-exist")
	if ok {
		t.Errorf("expected not to find transaction")
	}
}

func TestUpdateTransactionStatus_Success(t *testing.T) {
	repo := NewInMemoryTransactionRepository()
	tx := &models.Transaction{
		ID:      "tx2",
		Type:    constants.TypeWithdrawal,
		Amount:  50,
		Status:  constants.StatusPending,
		Account: "acc2",
	}
	repo.CreateTransaction(tx)

	err := repo.UpdateTransactionStatus("tx2", constants.StatusSuccess)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got, ok := repo.GetTransactionByID("tx2")
	if !ok {
		t.Fatalf("expected transaction to exist")
	}
	if got.Status != constants.StatusSuccess {
		t.Errorf("expected status %v, got %v", constants.StatusSuccess, got.Status)
	}
}

func TestUpdateTransactionStatus_NotFound(t *testing.T) {
	repo := NewInMemoryTransactionRepository()
	err := repo.UpdateTransactionStatus("not-exist", constants.StatusFailed)
	if err != nil {
		t.Errorf("expected no error when updating non-existent transaction, got %v", err)
	}
}

func TestCreateTransaction_Overwrite(t *testing.T) {
	repo := NewInMemoryTransactionRepository()
	tx := &models.Transaction{
		ID:      "tx3",
		Type:    constants.TypeDeposit,
		Amount:  10,
		Status:  constants.StatusPending,
		Account: "acc3",
	}
	repo.CreateTransaction(tx)

	// Overwrite with new data
	tx2 := &models.Transaction{
		ID:      "tx3",
		Type:    constants.TypeWithdrawal,
		Amount:  20,
		Status:  constants.StatusSuccess,
		Account: "acc3",
	}
	repo.CreateTransaction(tx2)

	got, ok := repo.GetTransactionByID("tx3")
	if !ok {
		t.Fatalf("expected transaction to exist")
	}
	if got.Type != constants.TypeWithdrawal || got.Amount != 20 || got.Status != constants.StatusSuccess {
		t.Errorf("transaction not overwritten as expected: %+v", got)
	}
}
