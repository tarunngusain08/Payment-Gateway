package handler

import (
	"Payment-Gateway/internal/service"
	"net/http"
)

func NewDepositHandler(svc *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Deposit handler implementation
	}
}
