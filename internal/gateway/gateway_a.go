package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/pkg/logger"

	"go.uber.org/zap"
)

// GatewayA is a skeleton for a JSON-based gateway.
type GatewayA struct {
	URL string
}

func NewGatewayA(url string) PaymentGateway {
	log := logger.GetLogger().With(zap.String("func", "NewGatewayA"))
	log.Info("Initializing GatewayA", zap.String("url", url))
	return &GatewayA{URL: url}
}

// ProcessDeposit simulates HTTP JSON request/response for GatewayA, handling success, failure, and timeout.
func (g *GatewayA) ProcessDeposit(r *http.Request) (interface{}, error) {
	log := logger.GetLogger().With(
		zap.String("func", "GatewayA.ProcessDeposit"),
		zap.String("url", g.URL),
	)
	var req dtos.GatewayADepositRequest
	if r != nil {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Warn("Failed to decode deposit request", zap.Error(err))
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
	log.Info("Sending deposit request to gateway")
	resp, err := client.Do(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Error("GatewayA deposit timeout", zap.Error(err))
			return nil, errors.New("gateway A timeout")
		}
		log.Error("GatewayA deposit error", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("GatewayA deposit failed", zap.Int("status_code", resp.StatusCode))
		return nil, errors.New("gateway A failure")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Error("Failed to decode gateway response", zap.Error(err))
		return nil, err
	}
	log.Info("GatewayA deposit successful", zap.Any("response", result))
	return result, nil
}

// ProcessWithdrawal simulates HTTP JSON request/response for GatewayA, handling success, failure, and timeout.
func (g *GatewayA) ProcessWithdrawal(r *http.Request) (interface{}, error) {
	log := logger.GetLogger().With(
		zap.String("func", "GatewayA.ProcessWithdrawal"),
		zap.String("url", g.URL),
	)
	var req dtos.GatewayAWithdrawalRequest
	if r != nil {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Warn("Failed to decode withdrawal request", zap.Error(err))
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
	log.Info("Sending withdrawal request to gateway")
	resp, err := client.Do(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Error("GatewayA withdrawal timeout", zap.Error(err))
			return nil, errors.New("gateway A timeout")
		}
		log.Error("GatewayA withdrawal error", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("GatewayA withdrawal failed", zap.Int("status_code", resp.StatusCode))
		return nil, errors.New("gateway A failure")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Error("Failed to decode gateway response", zap.Error(err))
		return nil, err
	}
	log.Info("GatewayA withdrawal successful", zap.Any("response", result))
	return result, nil
}
