package handler

import (
	"Payment-Gateway/internal/models"
	"Payment-Gateway/internal/service"
	"encoding/json"
	"net/http"
)

type TransactionRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type TransactionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type TransactionHandler struct {
	transactionService service.Transaction
}

func NewTransactionHandler(transactionService service.Transaction) TransactionHandler {
	return TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var req TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	depositReq := &models.DepositRequest{
		Account: req.AccountID,
		Amount:  req.Amount,
	}
	resp := TransactionResponse{}
	_, err := h.transactionService.CreateAndProcessDeposit(depositReq)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Success = true
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *TransactionHandler) Withdrawal(w http.ResponseWriter, r *http.Request) {
	var req TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	withdrawalReq := &models.WithdrawalRequest{
		Account: req.AccountID,
		Amount:  req.Amount,
	}
	resp := TransactionResponse{}
	_, err := h.transactionService.CreateAndProcessWithdrawal(withdrawalReq)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Success = true
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
