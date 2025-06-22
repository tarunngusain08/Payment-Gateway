package service

import (
	"Payment-Gateway/internal/gateway"
	errors "Payment-Gateway/pkg/error"
	"net/http"
	"testing"
)

type dummyGateway struct{}

func (d *dummyGateway) ProcessDeposit(r *http.Request) (interface{}, error)    { return nil, nil }
func (d *dummyGateway) ProcessWithdrawal(r *http.Request) (interface{}, error) { return nil, nil }

func TestGatewayPoolImpl_GetAllGateways(t *testing.T) {
	g1 := &dummyGateway{}
	g2 := &dummyGateway{}
	pool := NewGatewayPool([]gateway.PaymentGateway{g1, g2})

	gws, err := pool.GetAllGateways()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(gws) != 2 {
		t.Errorf("expected 2 gateways, got %d", len(gws))
	}
}

func TestGatewayPoolImpl_GetAllGateways_Empty(t *testing.T) {
	pool := NewGatewayPool([]gateway.PaymentGateway{})
	_, err := pool.GetAllGateways()
	if err != errors.ErrNoGatewayAvailable {
		t.Errorf("expected ErrNoGatewayAvailable, got %v", err)
	}
}

func TestGatewayPoolImpl_GetRoundRobinGateway(t *testing.T) {
	g1 := &dummyGateway{}
	g2 := &dummyGateway{}
	pool := NewGatewayPool([]gateway.PaymentGateway{g1, g2})

	gw1, _ := pool.GetRoundRobinGateway()
	gw2, _ := pool.GetRoundRobinGateway()
	gw3, _ := pool.GetRoundRobinGateway()

	if gw1 != g1 || gw2 != g2 || gw3 != g1 {
		t.Errorf("round robin logic failed")
	}
}

func TestGatewayPoolImpl_GetRoundRobinGateway_Empty(t *testing.T) {
	pool := NewGatewayPool([]gateway.PaymentGateway{})
	_, err := pool.GetRoundRobinGateway()
	if err != errors.ErrNoGatewayAvailable {
		t.Errorf("expected ErrNoGatewayAvailable, got %v", err)
	}
}
