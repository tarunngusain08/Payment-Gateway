package service

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/gateway"
)

// GatewayACallbackService handles callbacks from Gateway A
type GatewayACallbackService struct {
	// Add any dependencies or fields needed for the service
	gateway gateway.PaymentGateway
}

func NewGatewayACallbackService() Callback {
	return &GatewayACallbackService{
		gateway: gateway.GetGatewayByID("A"),
	}
}

// HandleCallback processes the callback request from Gateway A
// It returns a response and an HTTP status code
func (g GatewayACallbackService) HandleCallback(req dtos.HandleCallbackRequest) (dtos.HandleCallbackResponse, error) {
	// HandleCallback(w http.ResponseWriter, r *http.Request)
	g.gateway.HandleCallback()

	return response, nil
}
