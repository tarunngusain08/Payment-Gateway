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
)

// GatewayB is a skeleton for a SOAP-based gateway.
type GatewayB struct{}

func NewGatewayB() PaymentGateway {
	return &GatewayB{}
}

const gatewayBURL = "http://mock-gateway-b/soap" // Simulated endpoint

// ProcessDeposit simulates a SOAP request/response for GatewayB.
func (g *GatewayB) ProcessDeposit(r *http.Request) (interface{}, error) {
	req := &dtos.SOAPEnvelope{
		Body: dtos.SOAPBody{
			DepositRequest: &dtos.SOAPDepositRequest{Account: "demo", Amount: 100},
		},
	}
	payload, _ := xml.Marshal(req)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", gatewayBURL, bytes.NewBuffer(payload))
	httpReq.Header.Set("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("gateway B timeout")
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("gateway B failure")
	}

	body, _ := io.ReadAll(resp.Body)
	var envelope dtos.SOAPEnvelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}
	return envelope, nil
}

// ProcessWithdrawal simulates a SOAP request/response for GatewayB.
func (g *GatewayB) ProcessWithdrawal(r *http.Request) (interface{}, error) {
	req := &dtos.SOAPEnvelope{
		Body: dtos.SOAPBody{
			WithdrawalRequest: &dtos.SOAPWithdrawalRequest{Account: "demo", Amount: 100},
		},
	}
	payload, _ := xml.Marshal(req)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", gatewayBURL, bytes.NewBuffer(payload))
	httpReq.Header.Set("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("gateway B timeout")
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("gateway B failure")
	}

	body, _ := io.ReadAll(resp.Body)
	var envelope dtos.SOAPEnvelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}
	return envelope, nil
}
