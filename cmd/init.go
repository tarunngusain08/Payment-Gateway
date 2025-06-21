package main

import (
	"Payment-Gateway/internal/handler"
	"Payment-Gateway/internal/middleware"
	"Payment-Gateway/internal/repository"
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

func initializeHandlers() *handler.Handlers {
	repo := repository.NewInMemoryTransactionRepository()
	return &handler.Handlers{
		DepositHandler:    handler.NewDepositHandler(repo),
		WithdrawalHandler: handler.NewWithdrawalHandler(repo),
		GatewayACallback:  handler.NewGatewayACallback(repo),
		GatewayBCallback:  handler.NewGatewayBCallback(repo),
	}
}

func NewRouter() http.Handler {
	router := mux.NewRouter()
	handlers := initializeHandlers()

	initializeMiddlewares(router)

	// Setup routes
	setupRoutes(router, handlers)

	return router
}
