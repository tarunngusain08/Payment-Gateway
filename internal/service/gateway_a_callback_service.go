package service

import (
	"Payment-Gateway/internal/constants"
	"Payment-Gateway/internal/dtos"
	"fmt"
)

type GatewayACallbackService struct {
	transactionService Transaction
}

func NewGatewayACallbackService(transactionService Transaction) *GatewayACallbackService {
	return &GatewayACallbackService{
		transactionService: transactionService,
	}
}

func (g *GatewayACallbackService) HandleCallback(req dtos.HandleCallbackRequest) (*dtos.HandleCallbackResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	status := constants.TransactionStatus(req.Status)
	if err := g.transactionService.UpdateStatus(req.TransactionID, status); err != nil {
		return nil, fmt.Errorf("failed to update transaction status: %w", err)
	}

	return &dtos.HandleCallbackResponse{
		Status:  "success",
		Message: fmt.Sprintf("Successfully processed callback for transaction: %s", req.TransactionID),
	}, nil
}
