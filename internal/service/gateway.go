package service

import (
	"Payment-Gateway/internal/gateway"
	"sync"

	errors "Payment-Gateway/pkg/error"
	"Payment-Gateway/pkg/logger"

	"go.uber.org/zap"
)

type GatewayPoolImpl struct {
	gateways []gateway.PaymentGateway
	mu       sync.Mutex
	rrIndex  int
}

func NewGatewayPool(gateways []gateway.PaymentGateway) GatewayPool {
	log := logger.GetLogger().With(zap.String("func", "NewGatewayPool"))
	log.Info("Initializing GatewayPool", zap.Int("num_gateways", len(gateways)))
	return &GatewayPoolImpl{gateways: gateways}
}

func (gp *GatewayPoolImpl) GetAllGateways() ([]gateway.PaymentGateway, error) {
	log := logger.GetLogger().With(zap.String("func", "GatewayPoolImpl.GetAllGateways"))
	if len(gp.gateways) == 0 {
		log.Warn("No gateways available")
		return nil, errors.ErrNoGatewayAvailable
	}
	log.Info("Returning all gateways", zap.Int("count", len(gp.gateways)))
	return gp.gateways, nil
}

func (gp *GatewayPoolImpl) GetRoundRobinGateway() (gateway.PaymentGateway, error) {
	log := logger.GetLogger().With(zap.String("func", "GatewayPoolImpl.GetRoundRobinGateway"))
	gp.mu.Lock()
	defer gp.mu.Unlock()
	if len(gp.gateways) == 0 {
		log.Warn("No gateways available")
		return nil, errors.ErrNoGatewayAvailable
	}
	gateway := gp.gateways[gp.rrIndex%len(gp.gateways)]
	log.Info("Selected gateway", zap.Int("index", gp.rrIndex%len(gp.gateways)))
	gp.rrIndex = (gp.rrIndex + 1) % len(gp.gateways)
	return gateway, nil
}
