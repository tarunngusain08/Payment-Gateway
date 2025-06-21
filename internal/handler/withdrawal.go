package handler

import (
	"Payment-Gateway/internal/service"
	"net/http"
)

func NewWithdrawalHandler(svc *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Withdrawal handler implementation
	}
}
