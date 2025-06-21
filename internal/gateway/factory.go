package gateway

import (
	errors "Payment-Gateway/pkg/error"
)

// GetGatewayByID returns a PaymentGateway implementation based on the given ID.
func GetGatewayByID(id string) (PaymentGateway, error) {
	switch id {
	case "A":
		return &GatewayA{}, nil
	case "B":
		return &GatewayB{}, nil
	default:
		return nil, errors.ErrUnsupportedGateway
	}
}
