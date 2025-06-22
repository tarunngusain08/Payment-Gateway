package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type GatewayBCallbackService struct {
	transactionService Transaction
}

func NewGatewayBCallbackService(transactionService Transaction) Callback {
	return &GatewayBCallbackService{
		transactionService: transactionService,
	}
}

func (g *GatewayBCallbackService) HandleCallback(req dtos.HandleCallbackRequest) (*dtos.HandleCallbackResponse, error) {
	log := logger.GetLogger().With(
		zap.String("func", "GatewayBCallbackService.HandleCallback"),
		zap.String("transaction_id", req.TransactionID),
		zap.String("gateway_ref", req.GatewayRef),
		zap.String("status", req.Status),
	)
	log.Info("Handling GatewayB callback")

	if err := req.Validate(); err != nil {
		log.Warn("Validation failed", zap.Error(err))
		return nil, err
	}

	status := constants.TransactionStatus(req.Status)
	if err := g.transactionService.UpdateStatus(req.TransactionID, status); err != nil {
		log.Error("Failed to update transaction status", zap.Error(err))
		return nil, fmt.Errorf("failed to update transaction status: %w", err)
	}

	log.Info("GatewayB callback processed successfully")
	return &dtos.HandleCallbackResponse{
		Status: "success",
		Message: fmt.Sprintf("Gateway B callback processed for transaction: %s, ref: %s",
			req.TransactionID, req.GatewayRef),
	}, nil
}
