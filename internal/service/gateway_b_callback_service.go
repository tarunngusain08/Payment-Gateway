package service

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/gateway"
	errors "Payment-Gateway/pkg/error"
	"fmt"
)

type GatewayBCallbackService struct {
	gateway gateway.PaymentGateway
}

func NewGatewayBCallbackService() (Callback, error) {
	gw, err := gateway.GetGatewayByID("B")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gateway B: %w", err)
	}

	return &GatewayBCallbackService{
		gateway: gw,
	}, nil
}

func (g *GatewayBCallbackService) HandleCallback(req dtos.HandleCallbackRequest) (*dtos.HandleCallbackResponse, error) {
	// Validate required fields for Gateway B
	if req.TransactionID == "" || req.GatewayRef == "" {
		return nil, errors.ErrMissingRequiredFields
	}

	// Additional validation for Gateway B specific requirements
	if req.Amount <= 0 {
		return nil, errors.ErrMissingAmount
	}

	if req.Currency == "" {
		return nil, errors.ErrMissingCurrency
	}

	if err := g.gateway.HandleCallback(req); err != nil {
		return &dtos.HandleCallbackResponse{
			Status:  "failed",
			Message: fmt.Sprintf("Gateway B callback processing failed: %v", err),
		}, err
	}

	return &dtos.HandleCallbackResponse{
		Status: "success",
		Message: fmt.Sprintf("Gateway B callback processed for transaction: %s, ref: %s",
			req.TransactionID, req.GatewayRef),
	}, nil
}
