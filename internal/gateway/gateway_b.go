package gateway

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"time"

	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/pkg/logger"

	"Payment-Gateway/internal/config"

	"github.com/cenkalti/backoff/v4"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
)

// GatewayB is a skeleton for a SOAP-based gateway.
type GatewayB struct {
	URL            string
	Client         *http.Client
	CircuitBreaker *gobreaker.CircuitBreaker
	Config         config.ResilienceConfig
}

func NewGatewayB(url, gatewayName string) PaymentGateway {
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
	return &GatewayB{
		URL:            url,
		Client:         &http.Client{Timeout: time.Duration(cfg.HTTPTimeoutSeconds) * time.Second},
		CircuitBreaker: cb,
		Config:         cfg,
	}
}

func (g *GatewayB) doWithResilience(req *http.Request) (*http.Response, error) {
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

// ProcessDeposit simulates a SOAP request/response for GatewayB.
func (g *GatewayB) ProcessDeposit(r *http.Request) (interface{}, error) {
	log := logger.GetLogger().With(
		zap.String("func", "GatewayB.ProcessDeposit"),
		zap.String("url", g.URL),
	)
	req := &dtos.SOAPEnvelope{
		Body: dtos.SOAPBody{
			DepositRequest: &dtos.SOAPDepositRequest{Account: "demo", Amount: 100},
		},
	}
	payload, _ := xml.Marshal(req)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", g.URL+"/deposit", bytes.NewBuffer(payload))
	httpReq.Header.Set("Content-Type", "application/xml")

	log.Info("Sending deposit request to gateway")
	resp, err := g.doWithResilience(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Error("GatewayB deposit timeout", zap.Error(err))
			return nil, errors.New("gateway B timeout")
		}
		log.Error("GatewayB deposit error", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("GatewayB deposit failed", zap.Int("status_code", resp.StatusCode))
		return nil, errors.New("gateway B failure")
	}

	body, _ := io.ReadAll(resp.Body)
	var envelope dtos.SOAPEnvelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		log.Error("Failed to decode gateway response", zap.Error(err))
		return nil, err
	}
	log.Info("GatewayB deposit successful", zap.Any("response", envelope))
	return envelope, nil
}

// ProcessWithdrawal simulates a SOAP request/response for GatewayB.
func (g *GatewayB) ProcessWithdrawal(r *http.Request) (interface{}, error) {
	log := logger.GetLogger().With(
		zap.String("func", "GatewayB.ProcessWithdrawal"),
		zap.String("url", g.URL),
	)
	req := &dtos.SOAPEnvelope{
		Body: dtos.SOAPBody{
			WithdrawalRequest: &dtos.SOAPWithdrawalRequest{Account: "demo", Amount: 100},
		},
	}
	payload, _ := xml.Marshal(req)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", g.URL+"/withdrawal", bytes.NewBuffer(payload))
	httpReq.Header.Set("Content-Type", "application/xml")

	log.Info("Sending withdrawal request to gateway")
	resp, err := g.doWithResilience(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Error("GatewayB withdrawal timeout", zap.Error(err))
			return nil, errors.New("gateway B timeout")
		}
		log.Error("GatewayB withdrawal error", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("GatewayB withdrawal failed", zap.Int("status_code", resp.StatusCode))
		return nil, errors.New("gateway B failure")
	}

	body, _ := io.ReadAll(resp.Body)
	var envelope dtos.SOAPEnvelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		log.Error("Failed to decode gateway response", zap.Error(err))
		return nil, err
	}
	log.Info("GatewayB withdrawal successful", zap.Any("response", envelope))
	return envelope, nil
}
