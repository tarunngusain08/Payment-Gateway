package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// GatewayA is a skeleton for a JSON-based gateway.
type GatewayA struct{}

// DepositRequest represents a JSON deposit request for GatewayA.
type DepositRequest struct {
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

// WithdrawalRequest represents a JSON withdrawal request for GatewayA.
type WithdrawalRequest struct {
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

const gatewayAURL = "http://mock-gateway-a/deposit" // Simulated endpoint

// ProcessDeposit simulates HTTP JSON request/response for GatewayA, handling success, failure, and timeout.
func (g *GatewayA) ProcessDeposit(r *http.Request) (interface{}, error) {
	var req DepositRequest
	if r != nil {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}
	} else {
		req = DepositRequest{Account: "demo", Amount: 100}
	}

	payload, _ := json.Marshal(req)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", gatewayAURL, bytes.NewBuffer(payload))
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
	var req WithdrawalRequest
	if r != nil {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}
	} else {
		req = WithdrawalRequest{Account: "demo", Amount: 100}
	}

	payload, _ := json.Marshal(req)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", gatewayAURL, bytes.NewBuffer(payload))
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

// HandleCallback returns a JSON callback response for GatewayA.
func (g *GatewayA) HandleCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"gateway":"A","callback":"received"}`))
}
