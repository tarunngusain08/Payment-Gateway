package service

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/gateway"
	errors "Payment-Gateway/pkg/error"
	"fmt"
)

// GatewayACallbackService handles callbacks from Gateway A
type GatewayACallbackService struct {
	// Add any dependencies or fields needed for the service
	gateway gateway.PaymentGateway
}

func NewGatewayACallbackService() (Callback, error) {
	gw, err := gateway.GetGatewayByID("A")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gateway A: %w", err)
	}

	return &GatewayACallbackService{
		gateway: gw,
	}, nil
}

// HandleCallback processes the callback request from Gateway A
// It returns a response and an HTTP status code
func (g *GatewayACallbackService) HandleCallback(req dtos.HandleCallbackRequest) (*dtos.HandleCallbackResponse, error) {
	if req.TransactionID == "" {
		return nil, errors.ErrMissingTransactionID
	}

	if err := g.gateway.HandleCallback(req); err != nil {
		return &dtos.HandleCallbackResponse{
			Status:  "failed",
			Message: fmt.Sprintf("Failed to process callback: %v", err),
		}, err
	}

	return &dtos.HandleCallbackResponse{
		Status:  "success",
		Message: fmt.Sprintf("Successfully processed callback for transaction: %s", req.TransactionID),
	}, nil
}
