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

	"Payment-Gateway/internal/config"

	"github.com/cenkalti/backoff/v4"
	"github.com/sony/gobreaker"
)

type GatewayA struct {
	URL            string
	Client         *http.Client
	CircuitBreaker *gobreaker.CircuitBreaker
	Config         config.ResilienceConfig
}

func NewGatewayA(url, gatewayName string) PaymentGateway {
	cfg := config.GetConfig().Resilience
	var cb *gobreaker.CircuitBreaker
	if cfg.CircuitBreaker.Enabled {
		cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        gatewayName,
			MaxRequests: cfg.CircuitBreaker.MaxRequests,
			Interval:    time.Duration(cfg.CircuitBreaker.Interval) * time.Second,
			Timeout:     time.Duration(cfg.CircuitBreaker.Timeout) * time.Second,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.Requests >= 3 && failRatio >= cfg.CircuitBreaker.FailureRatio
			},
		})
	}
	return &GatewayA{
		URL:            url,
		Client:         &http.Client{Timeout: time.Duration(cfg.HTTPTimeoutSeconds) * time.Second},
		CircuitBreaker: cb,
		Config:         cfg,
	}
}

func (g *GatewayA) doWithResilience(req *http.Request) (*http.Response, error) {
	operation := func() (interface{}, error) {
		if g.CircuitBreaker != nil {
			resp, err := g.CircuitBreaker.Execute(func() (interface{}, error) {
				return g.Client.Do(req)
			})
			if err != nil {
				return nil, err
			}
			return resp.(*http.Response), nil
		}
		return g.Client.Do(req)
	}

	b := backoff.NewExponentialBackOff()
	b.InitialInterval = time.Duration(g.Config.InitialBackoffMillis) * time.Millisecond
	b.MaxInterval = time.Duration(g.Config.MaxBackoffMillis) * time.Millisecond
	b.MaxElapsedTime = time.Duration(g.Config.HTTPTimeoutSeconds*g.Config.MaxRetries) * time.Second

	var resp *http.Response
	err := backoff.Retry(func() error {
		r, err := operation()
		if err != nil {
			return err
		}
		resp = r.(*http.Response)
		return nil
	}, backoff.WithMaxRetries(b, uint64(g.Config.MaxRetries)))
	return resp, err
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

	log.Info("Sending deposit request to gateway")
	resp, err := g.doWithResilience(httpReq)
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

	log.Info("Sending withdrawal request to gateway")
	resp, err := g.doWithResilience(httpReq)
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
