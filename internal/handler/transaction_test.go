package handler

import (
	"Payment-Gateway/internal/models"
	errors "Payment-Gateway/pkg/error"
	"Payment-Gateway/pkg/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Payment-Gateway/internal/dtos"

	"github.com/golang/mock/gomock"
)

func TestTransactionHandler_Deposit_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().
		CreateAndProcessDeposit(gomock.Any()).
		Return(&models.Transaction{ID: "tx1"}, nil)

	handler := NewTransactionHandler(mockTx)
	reqBody := dtos.TransactionRequest{AccountID: "acc1", Amount: 100}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/deposit", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.Deposit(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestTransactionHandler_Deposit_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().
		CreateAndProcessDeposit(gomock.Any()).
		Return(nil, errors.ErrInvalidRequest)

	handler := NewTransactionHandler(mockTx)
	reqBody := dtos.TransactionRequest{AccountID: "acc1", Amount: 100}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/deposit", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.Deposit(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestTransactionHandler_Withdrawal_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().
		CreateAndProcessWithdrawal(gomock.Any()).
		Return(&models.Transaction{ID: "tx2"}, nil)

	handler := NewTransactionHandler(mockTx)
	reqBody := dtos.TransactionRequest{AccountID: "acc2", Amount: 50}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/withdrawal", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.Withdrawal(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestTransactionHandler_Withdrawal_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := mocks.NewMockTransaction(ctrl)
	mockTx.EXPECT().
		CreateAndProcessWithdrawal(gomock.Any()).
		Return(nil, errors.ErrInvalidRequest)

	handler := NewTransactionHandler(mockTx)
	reqBody := dtos.TransactionRequest{AccountID: "acc2", Amount: 50}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/withdrawal", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.Withdrawal(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}
