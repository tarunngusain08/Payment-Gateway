package main

import (
	"Payment-Gateway/internal/gateway"
	"Payment-Gateway/internal/handler"
	"Payment-Gateway/internal/middleware"
	"Payment-Gateway/internal/repository"
	"Payment-Gateway/internal/service"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func initializeMiddlewares(router *mux.Router) {
	router.Use(middleware.ContextMiddleware)
	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.TimeoutMiddleware(10 * time.Second))
	router.Use(middleware.LatencyTrackerMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)
}

func initializeHandlers() (*handler.Handlers, error) {
	transactionRepo := repository.NewInMemoryTransactionRepository()
	gatewayPool := service.NewGatewayPool(
		[]gateway.PaymentGateway{
			gateway.NewGatewayA(),
			gateway.NewGatewayB(),
		},
	)
	transactionService := service.NewTransactionService(transactionRepo, gatewayPool)

	gatewayACallbackService := service.NewGatewayACallbackService(transactionService)

	gatewayBCallbackService := service.NewGatewayBCallbackService(transactionService)

	return &handler.Handlers{
		TransactionHandler: handler.NewTransactionHandler(transactionService),
		GatewayACallback:   handler.NewGatewayACallback(gatewayACallbackService),
		GatewayBCallback:   handler.NewGatewayBCallback(gatewayBCallbackService),
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
