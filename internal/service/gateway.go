package service

import (
	"Payment-Gateway/internal/gateway"
	"sync"

	errors "Payment-Gateway/pkg/error"
)

type GatewayPoolImpl struct {
	gateways []gateway.PaymentGateway
	mu       sync.Mutex
	rrIndex  int
}

func NewGatewayPool(gateways []gateway.PaymentGateway) GatewayPool {
	return &GatewayPoolImpl{gateways: gateways}
}

func (gp *GatewayPoolImpl) GetAllGateways() ([]gateway.PaymentGateway, error) {
	if len(gp.gateways) == 0 {
		return nil, errors.ErrNoGatewayAvailable
	}
	return gp.gateways, nil
}

func (gp *GatewayPoolImpl) GetRoundRobinGateway() (gateway.PaymentGateway, error) {
	gp.mu.Lock()
	defer gp.mu.Unlock()
	if len(gp.gateways) == 0 {
		return nil, errors.ErrNoGatewayAvailable
	}
	gateway := gp.gateways[gp.rrIndex%len(gp.gateways)]
	gp.rrIndex = (gp.rrIndex + 1) % len(gp.gateways)
	return gateway, nil
}
