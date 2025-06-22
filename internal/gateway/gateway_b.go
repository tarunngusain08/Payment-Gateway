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

	"go.uber.org/zap"
)

// GatewayB is a skeleton for a SOAP-based gateway.
type GatewayB struct {
	URL string
}

func NewGatewayB(url string) PaymentGateway {
	log := logger.GetLogger().With(zap.String("func", "NewGatewayB"))
	log.Info("Initializing GatewayB", zap.String("url", url))
	return &GatewayB{URL: url}
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

	client := &http.Client{}
	log.Info("Sending deposit request to gateway")
	resp, err := client.Do(httpReq)
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

	client := &http.Client{}
	log.Info("Sending withdrawal request to gateway")
	resp, err := client.Do(httpReq)
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
