package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/dtos"
	"fmt"
)

type GatewayBCallbackService struct {
	transactionService Transaction
}

func NewGatewayBCallbackService(transactionService Transaction) Callback {
	return &GatewayACallbackService{
		transactionService: transactionService,
	}
}

func (g *GatewayBCallbackService) HandleCallback(req dtos.HandleCallbackRequest) (*dtos.HandleCallbackResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	status := constants.TransactionStatus(req.Status)
	if err := g.transactionService.UpdateStatus(req.TransactionID, status); err != nil {
		return nil, fmt.Errorf("failed to update transaction status: %w", err)
	}

	return &dtos.HandleCallbackResponse{
		Status: "success",
		Message: fmt.Sprintf("Gateway B callback processed for transaction: %s, ref: %s",
			req.TransactionID, req.GatewayRef),
	}, nil
}
