package main

import (
	"Payment-Gateway/internal/handler"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router, handlers *handler.Handlers) {
	// Payment routes
	router.HandleFunc("/deposit", handlers.TransactionHandler.Deposit).Methods("POST")
	router.HandleFunc("/withdrawal", handlers.TransactionHandler.Withdrawal).Methods("POST")

	// Callback routes
	router.HandleFunc("/callbacks/gateway-a", handlers.GatewayACallback.ServeHTTP).Methods("POST")
	router.HandleFunc("/callbacks/gateway-b", handlers.GatewayBCallback.ServeHTTP).Methods("POST")
}
