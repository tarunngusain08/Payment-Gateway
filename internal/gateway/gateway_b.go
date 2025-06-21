package gateway

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"time"
)

// GatewayB is a skeleton for a SOAP-based gateway.
type GatewayB struct{}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SOAPBody `xml:"Body"`
}

type SOAPBody struct {
	DepositRequest    *SOAPDepositRequest    `xml:"DepositRequest,omitempty"`
	WithdrawalRequest *SOAPWithdrawalRequest `xml:"WithdrawalRequest,omitempty"`
}

const gatewayBURL = "http://mock-gateway-b/soap" // Simulated endpoint

// SOAPDepositRequest represents a SOAP/XML deposit request for GatewayB.
type SOAPDepositRequest struct {
	XMLName xml.Name `xml:"DepositRequest"`
	Account string   `xml:"Account"`
	Amount  float64  `xml:"Amount"`
}

// SOAPWithdrawalRequest represents a SOAP/XML withdrawal request for GatewayB.
type SOAPWithdrawalRequest struct {
	XMLName xml.Name `xml:"WithdrawalRequest"`
	Account string   `xml:"Account"`
	Amount  float64  `xml:"Amount"`
}

// ProcessDeposit simulates a SOAP request/response for GatewayB.
func (g *GatewayB) ProcessDeposit(r *http.Request) (interface{}, error) {
	req := &SOAPEnvelope{
		Body: SOAPBody{
			DepositRequest: &SOAPDepositRequest{Account: "demo", Amount: 100},
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
	var envelope SOAPEnvelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}
	return envelope, nil
}

// ProcessWithdrawal simulates a SOAP request/response for GatewayB.
func (g *GatewayB) ProcessWithdrawal(r *http.Request) (interface{}, error) {
	req := &SOAPEnvelope{
		Body: SOAPBody{
			WithdrawalRequest: &SOAPWithdrawalRequest{Account: "demo", Amount: 100},
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
	var envelope SOAPEnvelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}
	return envelope, nil
}

// HandleCallback returns a SOAP-like XML callback response for GatewayB.
func (g *GatewayB) HandleCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(`<CallbackResponse><Gateway>B</Gateway><Callback>received</Callback></CallbackResponse>`))
}
