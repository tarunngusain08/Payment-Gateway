package main

import (
	"Payment-Gateway/internal/handler"
	"Payment-Gateway/internal/middleware"
	"Payment-Gateway/internal/repository"
	"Payment-Gateway/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func initializeMiddlewares(router *mux.Router) {
	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.TimeoutMiddleware(10 * time.Second))
	router.Use(middleware.LatencyTrackerMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)
}

func initializeHandlers() (*handler.Handlers, error) {
	transactionRepo := repository.NewInMemoryTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepo)

	gatewayACallback, err := service.NewGatewayACallbackService()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gateway A callback: %w", err)
	}

	gatewayBCallback, err := service.NewGatewayBCallbackService()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gateway B callback: %w", err)
	}

	return &handler.Handlers{
		DepositHandler:    handler.NewDepositHandler(transactionService),
		WithdrawalHandler: handler.NewWithdrawalHandler(transactionService),
		GatewayACallback:  handler.NewGatewayACallback(gatewayACallback),
		GatewayBCallback:  handler.NewGatewayBCallback(gatewayBCallback),
	}, nil
}

func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()
	handlers, err := initializeHandlers()
	if err != nil {
		return nil, err
	}

	initializeMiddlewares(router)
	setupRoutes(router, handlers)

	return router, nil
}
