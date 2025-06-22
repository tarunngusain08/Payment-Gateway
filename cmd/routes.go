package main

import (
	"Payment-Gateway/internal/handler"
	mockgateway "Payment-Gateway/internal/handler/mock_gateway"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router, handlers *handler.Handlers) {
	// Payment routes
	router.HandleFunc("/deposit", handlers.TransactionHandler.Deposit).Methods("POST")
	router.HandleFunc("/withdrawal", handlers.TransactionHandler.Withdrawal).Methods("POST")

	// Callback routes
	router.HandleFunc("/callbacks/gateway-a", handlers.GatewayACallback.ServeHTTP).Methods("POST")
	router.HandleFunc("/callbacks/gateway-b", handlers.GatewayBCallback.ServeHTTP).Methods("POST")

	// Mock gateway simulation routes (match config base + /deposit or /withdrawal)
	router.HandleFunc("/mock-gateway-a/deposit", mockgateway.GatewayAMockDepositHandler).Methods("POST")
	router.HandleFunc("/mock-gateway-a/withdrawal", mockgateway.GatewayAMockWithdrawalHandler).Methods("POST")
	router.HandleFunc("/mock-gateway-b/deposit", mockgateway.GatewayBMockDepositHandler).Methods("POST")
	router.HandleFunc("/mock-gateway-b/withdrawal", mockgateway.GatewayBMockWithdrawalHandler).Methods("POST")
}
