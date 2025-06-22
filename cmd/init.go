package main

import (
	"Payment-Gateway/internal/cache"
	cfg "Payment-Gateway/internal/config"
	"Payment-Gateway/internal/gateway"
	"Payment-Gateway/internal/handler"
	"Payment-Gateway/internal/middleware"
	"Payment-Gateway/internal/repository"
	"Payment-Gateway/internal/service"
	"Payment-Gateway/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Registry for gateway constructors
var gatewayRegistry = map[string]func(url, gatewayName string, cfg *cfg.ResilienceConfig) gateway.PaymentGateway{
	"gatewayA": gateway.NewGatewayA,
	"gatewayB": gateway.NewGatewayB,
	// Add more gateway constructors here as needed
}

func initializeMiddlewares(router *mux.Router) {
	router.Use(middleware.ContextMiddleware)
	router.Use(middleware.TimeoutMiddleware(10 * time.Second))
	router.Use(middleware.LatencyTrackerMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)
}

func initializeHandlers() (*handler.Handlers, error) {
	cfg := cfg.GetConfig()

	// Initialize cache with config values
	janitorInterval := time.Duration(cfg.Cache.InvalidationIntervalSeconds) * time.Second
	callbackCache := cache.NewMemoryCacheWithJanitor(janitorInterval, time.Duration(cfg.Cache.TTLSeconds)*time.Second)

	// Initialize worker pool
	numWorkers := cfg.WorkerPool.NumWorkers
	workerPool := service.NewWorkerPool(numWorkers)

	var gateways []gateway.PaymentGateway
	for name, gwCfg := range cfg.Gateways {
		if gwCfg.Enabled {
			if constructor, ok := gatewayRegistry[name]; ok {
				gateways = append(gateways, constructor(gwCfg.URL, gwCfg.Name, &cfg.Resilience))
			}
		}
	}

	transactionRepo := repository.NewInMemoryTransactionRepository()
	gatewayPool := service.NewGatewayPool(gateways)
	gatewayTimeout := time.Duration(cfg.Static.GatewayTimeoutSeconds) * time.Second
	transactionService := service.NewTransactionService(transactionRepo, gatewayPool, workerPool, gatewayTimeout)

	gatewayACallbackService := service.NewGatewayACallbackService(transactionService)
	gatewayBCallbackService := service.NewGatewayBCallbackService(transactionService)

	return &handler.Handlers{
		TransactionHandler: handler.NewTransactionHandler(transactionService),
		GatewayACallback:   handler.NewGatewayACallback(gatewayACallbackService, callbackCache),
		GatewayBCallback:   handler.NewGatewayBCallback(gatewayBCallbackService, callbackCache),
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

func StartServer() error {
	cfg := cfg.GetConfig()
	router, err := NewRouter()
	if err != nil {
		return err
	}
	addr := fmt.Sprintf("%s:%d", cfg.Static.Host, cfg.Static.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	defer logger.Sync()
	logger.GetLogger().Info("Starting server", zap.String("address", addr))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		logger.GetLogger().Info("Shutdown signal received")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.GetLogger().Error("Server shutdown error", zap.Error(err))
		}
	}()

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.GetLogger().Error("Server failed", zap.Error(err))
		return err
	}

	logger.GetLogger().Info("Server exited gracefully")
	return nil
}
