package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"Payment-Gateway/internal/dtos"
)

// GatewayA is a skeleton for a JSON-based gateway.
type GatewayA struct {
	URL string
}

func NewGatewayA(url string) PaymentGateway {
	return &GatewayA{URL: url}
}

// ProcessDeposit simulates HTTP JSON request/response for GatewayA, handling success, failure, and timeout.
func (g *GatewayA) ProcessDeposit(r *http.Request) (interface{}, error) {
	var req dtos.GatewayADepositRequest
	if r != nil {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}
	} else {
		req = dtos.GatewayADepositRequest{Account: "demo", Amount: 100}
	}

	payload, _ := json.Marshal(req)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", g.URL+"/deposit", bytes.NewBuffer(payload))
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("gateway A timeout")
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("gateway A failure")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ProcessWithdrawal simulates HTTP JSON request/response for GatewayA, handling success, failure, and timeout.
func (g *GatewayA) ProcessWithdrawal(r *http.Request) (interface{}, error) {
	var req dtos.GatewayAWithdrawalRequest
	if r != nil {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}
	} else {
		req = dtos.GatewayAWithdrawalRequest{Account: "demo", Amount: 100}
	}

	payload, _ := json.Marshal(req)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", g.URL+"/withdrawal", bytes.NewBuffer(payload))
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("gateway A timeout")
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("gateway A failure")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
