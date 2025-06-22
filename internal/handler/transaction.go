package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/middleware"
	"Payment-Gateway/internal/models"
	"Payment-Gateway/internal/service"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type TransactionHandler struct {
	transactionService service.Transaction
}

func NewTransactionHandler(transactionService service.Transaction) TransactionHandler {
	return TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(zap.String("func", "TransactionHandler.Deposit"))
	log.Info("Received deposit request")

	var req dtos.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Warn("Invalid deposit request payload", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log = log.With(zap.String("account_id", req.AccountID), zap.Float64("amount", req.Amount))
	depositReq := &models.DepositRequest{
		Account: req.AccountID,
		Amount:  req.Amount,
	}
	resp := dtos.TransactionResponse{}
	_, err := h.transactionService.CreateAndProcessDeposit(depositReq)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		log.Error("Deposit failed", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Success = true
		log.Info("Deposit successful")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *TransactionHandler) Withdrawal(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(zap.String("func", "TransactionHandler.Withdrawal"))
	log.Info("Received withdrawal request")

	var req dtos.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Warn("Invalid withdrawal request payload", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log = log.With(zap.String("account_id", req.AccountID), zap.Float64("amount", req.Amount))
	withdrawalReq := &models.WithdrawalRequest{
		Account: req.AccountID,
		Amount:  req.Amount,
	}
	resp := dtos.TransactionResponse{}
	_, err := h.transactionService.CreateAndProcessWithdrawal(withdrawalReq)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		log.Error("Withdrawal failed", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Success = true
		log.Info("Withdrawal successful")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
